package validate

// BoolTruef returns a Validator that fails if value is false.
// The error message is formatted using template and args.
func BoolTruef[T ~bool](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		if !bool(value) {
			return Errorf(template, args...)
		}
		return nil
	}
}

// BoolTrue returns a Validator that fails if value is false.
func BoolTrue[T ~bool](value T) ValidatorFunc {
	return BoolTruef(value, "must be true")
}

// BoolFalsef returns a Validator that fails if value is true.
// The error message is formatted using template and args.
func BoolFalsef[T ~bool](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		if bool(value) {
			return Errorf(template, args...)
		}
		return nil
	}
}

// BoolFalse returns a Validator that fails if value is true.
func BoolFalse[T ~bool](value T) ValidatorFunc {
	return BoolFalsef(value, "must be false")
}
