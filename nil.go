package validate

// NotNilf returns a Validator that fails if value is nil.
// If non-nil, it calls Validate() on the pointed-to value.
// The nil error message is formatted using template and args.
func NotNilf[T Validator](value *T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value == nil {
			return Errorf(template, args...)
		}
		return (*value).Validate()
	}
}

// NotNil returns a Validator that fails if value is nil.
// If non-nil, it calls Validate() on the pointed-to value.
func NotNil[T Validator](value *T) ValidatorFunc {
	return NotNilf(value, "is required")
}

// IfNotNil returns a Validator that passes when value is nil.
// If non-nil, it calls Validate() on the pointed-to value.
func IfNotNil[T Validator](value *T) ValidatorFunc {
	return func() error {
		if value == nil {
			return nil
		}
		return (*value).Validate()
	}
}
