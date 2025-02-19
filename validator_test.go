package validate

import (
	"errors"
	"testing"

	"github.com/kyuff/validate/internal/assert"
)

func TestValidatorFunc_Validate(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		// arrange
		var (
			sut = ValidatorFunc(func() error {
				return nil
			})
		)

		// act
		err := sut.Validate()

		// assert
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		// arrange
		var (
			sut = ValidatorFunc(func() error {
				return errors.New("test")
			})
		)

		// act
		err := sut.Validate()

		// assert
		assert.Error(t, err)
	})
}
