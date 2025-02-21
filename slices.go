package validate

import "slices"

func SliceContainsf[T comparable](values []T, target T, template string, args ...any) error {
	if !slices.Contains(values, target) {
		return Errorf(template, args...)
	}

	return nil
}
