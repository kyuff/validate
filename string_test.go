package validate_test

import (
	"regexp"
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

func TestStrings(t *testing.T) {
	t.Run("StringNotEmpty", func(t *testing.T) {
		t.Run("fail on empty string", func(t *testing.T) {
			// arrange
			var (
				value = ""
			)

			// act
			got := validate.StringNotEmpty(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass on non-empty string", func(t *testing.T) {
			// arrange
			var (
				value = "hello"
			)

			// act
			got := validate.StringNotEmpty(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("StringNotEmptyf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = ""
			)

			// act
			got := validate.StringNotEmptyf(value, "id is required").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "id is required", got.Error())
		})

		t.Run("pass on non-empty string", func(t *testing.T) {
			// arrange
			var (
				value = "hello"
			)

			// act
			got := validate.StringNotEmptyf(value, "id is required").Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("StringMinLength", func(t *testing.T) {
		t.Run("fail below minimum length", func(t *testing.T) {
			// arrange
			var (
				value = "hi"
				min   = 5
			)

			// act
			got := validate.StringMinLength(value, min).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass at exact minimum length", func(t *testing.T) {
			// arrange
			var (
				value = "abc"
				min   = len(value)
			)

			// act
			got := validate.StringMinLength(value, min).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass above minimum length", func(t *testing.T) {
			// arrange
			var (
				value = "hello"
				min   = 3
			)

			// act
			got := validate.StringMinLength(value, min).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("StringMinLengthf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = "hi"
			)

			// act
			got := validate.StringMinLengthf(value, 5, "too short").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "too short", got.Error())
		})
	})

	t.Run("StringMaxLength", func(t *testing.T) {
		t.Run("fail above maximum length", func(t *testing.T) {
			// arrange
			var (
				value = "hello world"
				max   = 5
			)

			// act
			got := validate.StringMaxLength(value, max).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass at exact maximum length", func(t *testing.T) {
			// arrange
			var (
				value = "abc"
				max   = len(value)
			)

			// act
			got := validate.StringMaxLength(value, max).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass below maximum length", func(t *testing.T) {
			// arrange
			var (
				value = "hello"
				max   = 10
			)

			// act
			got := validate.StringMaxLength(value, max).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("StringMaxLengthf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = "hello world"
			)

			// act
			got := validate.StringMaxLengthf(value, 5, "too long").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "too long", got.Error())
		})
	})

	t.Run("StringMatches", func(t *testing.T) {
		t.Run("fail when pattern does not match", func(t *testing.T) {
			// arrange
			var (
				value   = "HELLO"
				pattern = regexp.MustCompile(`^[a-z]+$`)
			)

			// act
			got := validate.StringMatches(value, pattern).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass when pattern matches", func(t *testing.T) {
			// arrange
			var (
				value   = "hello123"
				pattern = regexp.MustCompile(`^[a-z0-9]+$`)
			)

			// act
			got := validate.StringMatches(value, pattern).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("StringMatchesf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value   = "HELLO"
				pattern = regexp.MustCompile(`^[a-z]+$`)
			)

			// act
			got := validate.StringMatchesf(value, pattern, "invalid format").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "invalid format", got.Error())
		})
	})

	t.Run("StringOneOf", func(t *testing.T) {
		t.Run("fail when value not in options", func(t *testing.T) {
			// arrange
			var (
				value   = "deleted"
				options = []string{"active", "inactive"}
			)

			// act
			got := validate.StringOneOf(value, options...).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass when value in options", func(t *testing.T) {
			// arrange
			var (
				value   = "active"
				options = []string{"active", "inactive", "pending"}
			)

			// act
			got := validate.StringOneOf(value, options...).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("StringOneOff", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value   = "deleted"
				options = []string{"active", "inactive"}
			)

			// act
			got := validate.StringOneOff(value, options, "unknown status").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "unknown status", got.Error())
		})

		t.Run("pass when value in options", func(t *testing.T) {
			// arrange
			var (
				value   = "active"
				options = []string{"active", "inactive"}
			)

			// act
			got := validate.StringOneOff(value, options, "unknown status").Validate()

			// assert
			assert.NoError(t, got)
		})
	})
}
