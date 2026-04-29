package validate

import (
	"regexp"
	"slices"
)

// StringNotEmptyf returns a Validator that fails if value is an empty string.
// The error message is formatted using template and args.
func StringNotEmptyf[T ~string](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value == "" {
			return Errorf(template, args...)
		}
		return nil
	}
}

// StringNotEmpty returns a Validator that fails if value is an empty string.
func StringNotEmpty[T ~string](value T) ValidatorFunc {
	return StringNotEmptyf(value, "must not be empty")
}

// StringMinLengthf returns a Validator that fails if value has fewer than min bytes.
// The error message is formatted using template and args.
func StringMinLengthf[T ~string](value T, min int, template string, args ...any) ValidatorFunc {
	return func() error {
		if len(value) < min {
			return Errorf(template, args...)
		}
		return nil
	}
}

// StringMinLength returns a Validator that fails if value has fewer than min bytes.
func StringMinLength[T ~string](value T, min int) ValidatorFunc {
	return StringMinLengthf(value, min, "must be at least %d characters", min)
}

// StringMaxLengthf returns a Validator that fails if value exceeds max bytes.
// The error message is formatted using template and args.
func StringMaxLengthf[T ~string](value T, max int, template string, args ...any) ValidatorFunc {
	return func() error {
		if len(value) > max {
			return Errorf(template, args...)
		}
		return nil
	}
}

// StringMaxLength returns a Validator that fails if value exceeds max bytes.
func StringMaxLength[T ~string](value T, max int) ValidatorFunc {
	return StringMaxLengthf(value, max, "must be at most %d characters", max)
}

// StringMatchesf returns a Validator that fails if value does not match pattern.
// The error message is formatted using template and args.
func StringMatchesf[T ~string](value T, pattern *regexp.Regexp, template string, args ...any) ValidatorFunc {
	return func() error {
		if !pattern.MatchString(string(value)) {
			return Errorf(template, args...)
		}
		return nil
	}
}

// StringMatches returns a Validator that fails if value does not match pattern.
func StringMatches[T ~string](value T, pattern *regexp.Regexp) ValidatorFunc {
	return StringMatchesf(value, pattern, "must match %v", pattern)
}

// StringOneOff returns a Validator that fails if value is not present in options.
// The error message is formatted using template and args.
func StringOneOff[T ~string](value T, options []T, template string, args ...any) ValidatorFunc {
	return func() error {
		if slices.Contains(options, value) {
			return nil
		}
		return Errorf(template, args...)
	}
}

// StringOneOf returns a Validator that fails if value is not present in options.
func StringOneOf[T ~string](value T, options ...T) ValidatorFunc {
	return StringOneOff(value, options, "%v must be one of %v", value, options)
}
