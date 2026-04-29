package validate

// Requiredf returns a Validator that fails if value equals its zero value.
// Works with any comparable type, including named types such as type ID string.
// The error message is formatted using template and args.
func Requiredf[T comparable](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		var zero T
		if value == zero {
			return Errorf(template, args...)
		}
		return nil
	}
}

// Required returns a Validator that fails if value equals its zero value.
// Works with any comparable type, including named types such as type ID string.
func Required[T comparable](value T) ValidatorFunc {
	return Requiredf(value, "is required")
}
