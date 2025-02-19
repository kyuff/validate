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
			return Errorf("not a validator: %T", arg)
		}

		if isNil(reflect.ValueOf(v)) {
			return Errorf("nil: %T", arg)
		}

		err := v.Validate()
		if err != nil {
			return Error{err: fmt.Errorf("invalid: %w", err)}
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
