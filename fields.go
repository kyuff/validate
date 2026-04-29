package validate

import (
	"errors"
	"reflect"
	"sync"
)

var validatorType = reflect.TypeOf((*Validator)(nil)).Elem()

type fieldsStrictEntry struct {
	err     error
	indices []int
}

var fieldsStrictCache sync.Map // map[reflect.Type]*fieldsStrictEntry

func loadFieldsStrictEntry(t reflect.Type) *fieldsStrictEntry {
	if v, ok := fieldsStrictCache.Load(t); ok {
		return v.(*fieldsStrictEntry)
	}

	entry := buildFieldsStrictEntry(t)
	actual, _ := fieldsStrictCache.LoadOrStore(t, entry)
	return actual.(*fieldsStrictEntry)
}

func buildFieldsStrictEntry(t reflect.Type) *fieldsStrictEntry {
	indices := make([]int, 0, t.NumField())
	for i := range t.NumField() {
		field := t.Field(i)
		if !field.Type.Implements(validatorType) {
			return &fieldsStrictEntry{
				err: Errorf("field %s does not implement Validator", field.Name),
			}
		}
		indices = append(indices, i)
	}
	return &fieldsStrictEntry{indices: indices}
}

// FieldsStrict returns a Validator that iterates all fields of a struct and
// calls Validate() on each one. It fails immediately if any field does not
// implement the Validator interface, and collects validation errors from all fields.
// Accepts a struct or a pointer to a struct.
// Reflection metadata is cached per type, so repeated calls for the same type
// avoid redundant reflection overhead.
func FieldsStrict(s any) ValidatorFunc {
	return func() error {
		v := reflect.ValueOf(s)

		for v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return Errorf("must not be nil")
			}
			v = v.Elem()
		}

		if v.Kind() != reflect.Struct {
			return Errorf("must be a struct")
		}

		entry := loadFieldsStrictEntry(v.Type())
		if entry.err != nil {
			return entry.err
		}

		var err error
		for _, i := range entry.indices {
			validator := v.Field(i).Interface().(Validator)
			err = errors.Join(err, validator.Validate())
		}

		return err
	}
}
