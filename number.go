package validate

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type float interface {
	~float32 | ~float64
}

type number interface {
	integer | float
}

// NumberMinf returns a Validator that fails if value is less than min.
// The error message is formatted using template and args.
func NumberMinf[T number](value, min T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value < min {
			return Errorf(template, args...)
		}
		return nil
	}
}

// NumberMin returns a Validator that fails if value is less than min.
func NumberMin[T number](value, min T) ValidatorFunc {
	return NumberMinf(value, min, "must be at least %v", min)
}

// NumberMaxf returns a Validator that fails if value is greater than max.
// The error message is formatted using template and args.
func NumberMaxf[T number](value, max T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value > max {
			return Errorf(template, args...)
		}
		return nil
	}
}

// NumberMax returns a Validator that fails if value is greater than max.
func NumberMax[T number](value, max T) ValidatorFunc {
	return NumberMaxf(value, max, "must be at most %v", max)
}

// NumberPositivef returns a Validator that fails if value is zero or negative.
// The error message is formatted using template and args.
func NumberPositivef[T number](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value <= 0 {
			return Errorf(template, args...)
		}
		return nil
	}
}

// NumberPositive returns a Validator that fails if value is zero or negative.
func NumberPositive[T number](value T) ValidatorFunc {
	return NumberPositivef(value, "must be positive")
}

// NumberNegativef returns a Validator that fails if value is zero or positive.
// The error message is formatted using template and args.
func NumberNegativef[T number](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value >= 0 {
			return Errorf(template, args...)
		}
		return nil
	}
}

// NumberNegative returns a Validator that fails if value is zero or positive.
func NumberNegative[T number](value T) ValidatorFunc {
	return NumberNegativef(value, "must be negative")
}

// NumberBetweenf returns a Validator that fails if value is outside the inclusive range [min, max].
// The error message is formatted using template and args.
func NumberBetweenf[T number](value, min, max T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value < min || value > max {
			return Errorf(template, args...)
		}
		return nil
	}
}

// NumberBetween returns a Validator that fails if value is outside the inclusive range [min, max].
func NumberBetween[T number](value, min, max T) ValidatorFunc {
	return NumberBetweenf(value, min, max, "must be between %v and %v", min, max)
}
