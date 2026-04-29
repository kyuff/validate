package validate_test

import (
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

func TestRequired(t *testing.T) {
	t.Run("Required", func(t *testing.T) {
		t.Run("fail on zero string", func(t *testing.T) {
			// arrange
			var (
				value = ""
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("fail on zero int", func(t *testing.T) {
			// arrange
			var (
				value = 0
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("fail on zero bool", func(t *testing.T) {
			// arrange
			var (
				value = false
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass on non-zero string", func(t *testing.T) {
			// arrange
			var (
				value = "hello"
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass on non-zero int", func(t *testing.T) {
			// arrange
			var (
				value = 42
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass with named string type", func(t *testing.T) {
			// arrange
			type ID string
			var (
				value = ID("abc")
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("fail with named string type at zero value", func(t *testing.T) {
			// arrange
			type ID string
			var (
				value = ID("")
			)

			// act
			got := validate.Required(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})
	})

	t.Run("Requiredf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = ""
			)

			// act
			got := validate.Requiredf(value, "user id is required").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "user id is required", got.Error())
		})
	})
}
