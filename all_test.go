package validate_test

import (
	"errors"
	"testing"

	"github.com/kyuff/validate"
)

func TestAll(t *testing.T) {
	t.Run("ignore nil validators", func(t *testing.T) {
		// act
		err := validate.All(nil)

		// assert
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("fail with validator", func(t *testing.T) {
		// arrange
		var (
			v = &ValidatorMock{
				ValidateFunc: func() error {
					return errors.New("test")
				},
			}
		)
		// act
		err := validate.All(v)

		// assert
		if err == nil {
			t.Fatal(err)
		}
	})
	t.Run("validate multiple", func(t *testing.T) {
		// arrange
		var (
			castTo = func(ms []*ValidatorMock) []validate.Validator {
				var result []validate.Validator
				for _, m := range ms {
					result = append(result, m)
				}
				return result
			}
			inputs = []*ValidatorMock{
				{ValidateFunc: func() error {
					return nil
				}},
				{ValidateFunc: func() error {
					return errors.New("test")
				}},
				{ValidateFunc: func() error {
					return nil
				}},
			}
		)
		// act
		err := validate.All(castTo(inputs)...)

		// assert
		if err == nil {
			t.Fatal(err)
		}
	})
}
