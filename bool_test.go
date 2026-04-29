package validate_test

import (
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

func TestBools(t *testing.T) {
	t.Run("BoolTrue", func(t *testing.T) {
		t.Run("fail on false value", func(t *testing.T) {
			// arrange
			var (
				value = false
			)

			// act
			got := validate.BoolTrue(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass on true value", func(t *testing.T) {
			// arrange
			var (
				value = true
			)

			// act
			got := validate.BoolTrue(value).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass with named bool type", func(t *testing.T) {
			// arrange
			type Flag bool
			var (
				value = Flag(true)
			)

			// act
			got := validate.BoolTrue(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("BoolTruef", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = false
			)

			// act
			got := validate.BoolTruef(value, "must be accepted").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "must be accepted", got.Error())
		})
	})

	t.Run("BoolFalse", func(t *testing.T) {
		t.Run("fail on true value", func(t *testing.T) {
			// arrange
			var (
				value = true
			)

			// act
			got := validate.BoolFalse(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass on false value", func(t *testing.T) {
			// arrange
			var (
				value = false
			)

			// act
			got := validate.BoolFalse(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("BoolFalsef", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = true
			)

			// act
			got := validate.BoolFalsef(value, "must not be set").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "must not be set", got.Error())
		})
	})
}
