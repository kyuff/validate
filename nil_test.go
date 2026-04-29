package validate_test

import (
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

type nilTestID string

func (id nilTestID) Validate() error {
	return validate.StringNotEmpty(id).Validate()
}

func TestNil(t *testing.T) {
	newID := func(s string) *nilTestID { v := nilTestID(s); return &v }

	t.Run("NotNil", func(t *testing.T) {
		t.Run("fail on nil", func(t *testing.T) {
			// arrange
			var (
				value *nilTestID = nil
			)

			// act
			got := validate.NotNil(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("fail when pointed-to value is invalid", func(t *testing.T) {
			// arrange
			var (
				value = newID("")
			)

			// act
			got := validate.NotNil(value).Validate()

			// assert
			assert.Error(t, got)
		})

		t.Run("pass when non-nil and valid", func(t *testing.T) {
			// arrange
			var (
				value = newID("abc")
			)

			// act
			got := validate.NotNil(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("NotNilf", func(t *testing.T) {
		t.Run("fail with custom message on nil", func(t *testing.T) {
			// arrange
			var (
				value *nilTestID = nil
			)

			// act
			got := validate.NotNilf(value, "user is required").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "user is required", got.Error())
		})
	})

	t.Run("IfNotNil", func(t *testing.T) {
		t.Run("pass on nil", func(t *testing.T) {
			// arrange
			var (
				value *nilTestID = nil
			)

			// act
			got := validate.IfNotNil(value).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("fail when pointed-to value is invalid", func(t *testing.T) {
			// arrange
			var (
				value = newID("")
			)

			// act
			got := validate.IfNotNil(value).Validate()

			// assert
			assert.Error(t, got)
		})

		t.Run("pass when non-nil and valid", func(t *testing.T) {
			// arrange
			var (
				value = newID("abc")
			)

			// act
			got := validate.IfNotNil(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})
}
