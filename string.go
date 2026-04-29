package validate

import "regexp"

func StringNotEmptyf[T ~string](value T, template string, args ...any) ValidatorFunc {
	return func() error {
		if value == "" {
			return Errorf(template, args...)
		}
		return nil
	}
}

func StringNotEmpty[T ~string](value T) ValidatorFunc {
	return StringNotEmptyf(value, "must not be empty")
}

func StringMinLengthf[T ~string](value T, min int, template string, args ...any) ValidatorFunc {
	return func() error {
		if len(value) < min {
			return Errorf(template, args...)
		}
		return nil
	}
}

func StringMinLength[T ~string](value T, min int) ValidatorFunc {
	return StringMinLengthf(value, min, "must be at least %d characters", min)
}

func StringMaxLengthf[T ~string](value T, max int, template string, args ...any) ValidatorFunc {
	return func() error {
		if len(value) > max {
			return Errorf(template, args...)
		}
		return nil
	}
}

func StringMaxLength[T ~string](value T, max int) ValidatorFunc {
	return StringMaxLengthf(value, max, "must be at most %d characters", max)
}

func StringMatchesf[T ~string](value T, pattern *regexp.Regexp, template string, args ...any) ValidatorFunc {
	return func() error {
		if !pattern.MatchString(string(value)) {
			return Errorf(template, args...)
		}
		return nil
	}
}

func StringMatches[T ~string](value T, pattern *regexp.Regexp) ValidatorFunc {
	return StringMatchesf(value, pattern, "must match %v", pattern)
}

func StringOneOff[T ~string](value T, options []T, template string, args ...any) ValidatorFunc {
	return func() error {
		for _, o := range options {
			if value == o {
				return nil
			}
		}
		return Errorf(template, args...)
	}
}

func StringOneOf[T ~string](value T, options ...T) ValidatorFunc {
	return StringOneOff(value, options, "%v must be one of %v", value, options)
}
