package validate

import (
	"context"
	"fmt"
	"reflect"
)

func Middleware[T any](next func(ctx context.Context, arg T) error) func(ctx context.Context, arg T) error {
	return func(ctx context.Context, arg T) error {
		v, ok := any(arg).(Validator)
		if !ok {
			return fmt.Errorf("not a validator")
		}

		if isNil(reflect.ValueOf(v)) {
			return fmt.Errorf("arg is nil")
		}

		err := v.Validate()
		if err != nil {
			return err
		}

		return next(ctx, arg)
	}
}

func isNil(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return value.IsNil()
	default:
		return false
	}
}
