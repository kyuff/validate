package validate_test

import (
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

func TestSliceContainsf(t *testing.T) {
	t.Run("slice contains success", func(t *testing.T) {
		// act
		err := validate.SliceContainsf([]int{1, 2, 3}, 2, "must contain")

		// assert
		assert.NoError(t, err)
	})

	t.Run("slice contains fail", func(t *testing.T) {
		// act
		err := validate.SliceContainsf([]int{1, 2, 3}, 5, "must not contain")

		// assert
		assert.Error(t, err)
		assert.ErrorIs(t, validate.Error{}, err)
	})
}
