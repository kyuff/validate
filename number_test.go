package validate_test

import (
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

func TestNumbers(t *testing.T) {
	t.Run("NumberMin", func(t *testing.T) {
		t.Run("fail below minimum", func(t *testing.T) {
			// arrange
			var (
				value = 3
				min   = 5
			)

			// act
			got := validate.NumberMin(value, min).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass at exact minimum", func(t *testing.T) {
			// arrange
			var (
				value = 5
				min   = 5
			)

			// act
			got := validate.NumberMin(value, min).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass above minimum", func(t *testing.T) {
			// arrange
			var (
				value = 10
				min   = 5
			)

			// act
			got := validate.NumberMin(value, min).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass with float type", func(t *testing.T) {
			// arrange
			var (
				value = 3.14
				min   = 2.0
			)

			// act
			got := validate.NumberMin(value, min).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass with named int type", func(t *testing.T) {
			// arrange
			type Count int
			var (
				value = Count(10)
				min   = Count(5)
			)

			// act
			got := validate.NumberMin(value, min).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("NumberMinf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = 3
			)

			// act
			got := validate.NumberMinf(value, 5, "count too low").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "count too low", got.Error())
		})
	})

	t.Run("NumberMax", func(t *testing.T) {
		t.Run("fail above maximum", func(t *testing.T) {
			// arrange
			var (
				value = 15
				max   = 10
			)

			// act
			got := validate.NumberMax(value, max).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass at exact maximum", func(t *testing.T) {
			// arrange
			var (
				value = 10
				max   = 10
			)

			// act
			got := validate.NumberMax(value, max).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass below maximum", func(t *testing.T) {
			// arrange
			var (
				value = 5
				max   = 10
			)

			// act
			got := validate.NumberMax(value, max).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("NumberMaxf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = 15
			)

			// act
			got := validate.NumberMaxf(value, 10, "count too high").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "count too high", got.Error())
		})
	})

	t.Run("NumberPositive", func(t *testing.T) {
		t.Run("fail on zero", func(t *testing.T) {
			// arrange
			var (
				value = 0
			)

			// act
			got := validate.NumberPositive(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("fail on negative value", func(t *testing.T) {
			// arrange
			var (
				value = -1
			)

			// act
			got := validate.NumberPositive(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass on positive value", func(t *testing.T) {
			// arrange
			var (
				value = 5
			)

			// act
			got := validate.NumberPositive(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("NumberPositivef", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = 0
			)

			// act
			got := validate.NumberPositivef(value, "must be greater than zero").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "must be greater than zero", got.Error())
		})
	})

	t.Run("NumberNegative", func(t *testing.T) {
		t.Run("fail on zero", func(t *testing.T) {
			// arrange
			var (
				value = 0
			)

			// act
			got := validate.NumberNegative(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("fail on positive value", func(t *testing.T) {
			// arrange
			var (
				value = 1
			)

			// act
			got := validate.NumberNegative(value).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass on negative value", func(t *testing.T) {
			// arrange
			var (
				value = -5
			)

			// act
			got := validate.NumberNegative(value).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("NumberNegativef", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = 0
			)

			// act
			got := validate.NumberNegativef(value, "must be less than zero").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "must be less than zero", got.Error())
		})
	})

	t.Run("NumberBetween", func(t *testing.T) {
		t.Run("fail below range", func(t *testing.T) {
			// arrange
			var (
				value = 0
				min   = 1
				max   = 10
			)

			// act
			got := validate.NumberBetween(value, min, max).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("fail above range", func(t *testing.T) {
			// arrange
			var (
				value = 11
				min   = 1
				max   = 10
			)

			// act
			got := validate.NumberBetween(value, min, max).Validate()

			// assert
			assert.Error(t, got)
			assert.ErrorIs(t, validate.Error{}, got)
		})

		t.Run("pass at minimum boundary", func(t *testing.T) {
			// arrange
			var (
				value = 1
				min   = 1
				max   = 10
			)

			// act
			got := validate.NumberBetween(value, min, max).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass at maximum boundary", func(t *testing.T) {
			// arrange
			var (
				value = 10
				min   = 1
				max   = 10
			)

			// act
			got := validate.NumberBetween(value, min, max).Validate()

			// assert
			assert.NoError(t, got)
		})

		t.Run("pass within range", func(t *testing.T) {
			// arrange
			var (
				value = 5
				min   = 1
				max   = 10
			)

			// act
			got := validate.NumberBetween(value, min, max).Validate()

			// assert
			assert.NoError(t, got)
		})
	})

	t.Run("NumberBetweenf", func(t *testing.T) {
		t.Run("fail with custom message", func(t *testing.T) {
			// arrange
			var (
				value = 0
			)

			// act
			got := validate.NumberBetweenf(value, 1, 10, "out of range").Validate()

			// assert
			assert.Error(t, got)
			assert.Match(t, "out of range", got.Error())
		})
	})
}
