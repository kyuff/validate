package validate_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

type TestArg struct {
	ValidateFunc func() error
}

func (arg TestArg) Validate() error {
	return arg.ValidateFunc()
}

func TestMiddleware2(t *testing.T) {
	t.Run("fail on not implementing Validator", func(t *testing.T) {
		// arrange
		var (
			sut = validate.Middleware(func(ctx context.Context, arg string) error {
				return nil
			})
		)

		// act
		err := sut(t.Context(), "not a Validator")

		// assert
		assert.Error(t, err)
	})

	t.Run("fail on nil Validator", func(t *testing.T) {
		// arrange
		var (
			arg *ValidatorMock
			sut = validate.Middleware(func(ctx context.Context, arg *ValidatorMock) error {
				return nil
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.Error(t, err)
	})

	t.Run("fail on invalid Validator", func(t *testing.T) {
		// arrange
		var (
			arg = &ValidatorMock{
				ValidateFunc: func() error {
					return errors.New("invalid validator")
				},
			}
			sut = validate.Middleware(func(ctx context.Context, arg *ValidatorMock) error {
				return nil
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.Error(t, err)
		assert.Equal(t, 1, len(arg.ValidateCalls()))
	})

	t.Run("fail with next on valid Validator", func(t *testing.T) {
		// arrange
		var (
			called = false
			arg    = &ValidatorMock{
				ValidateFunc: func() error {
					return errors.New("invalid validator")
				},
			}
			sut = validate.Middleware(func(ctx context.Context, arg *ValidatorMock) error {
				called = true
				return errors.New("invalid next")
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.Error(t, err)
		assert.Truef(t, !called, "called")
		assert.Equal(t, 1, len(arg.ValidateCalls()))
	})

	t.Run("call next on valid pointer Validator", func(t *testing.T) {
		// arrange
		var (
			called = false
			arg    = &ValidatorMock{
				ValidateFunc: func() error {
					return nil
				},
			}
			sut = validate.Middleware(func(ctx context.Context, arg *ValidatorMock) error {
				called = true
				return nil
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.NoError(t, err)
		assert.Truef(t, called, "called")
		assert.Equal(t, 1, len(arg.ValidateCalls()))
	})

	t.Run("call next on valid value Validator", func(t *testing.T) {
		// arrange
		var (
			called = false
			arg    = TestArg{ValidateFunc: func() error {
				return nil
			}}
			sut = validate.Middleware(func(ctx context.Context, arg TestArg) error {
				called = true
				return nil
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.NoError(t, err)
		assert.Truef(t, called, "called")
	})

	t.Run("upgrade error to validate.Error", func(t *testing.T) {
		// arrange
		var (
			called = false
			arg    = TestArg{ValidateFunc: func() error {
				return errors.New("validation error")
			}}
			sut = validate.Middleware(func(ctx context.Context, arg TestArg) error {
				called = true
				return nil
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.Error(t, err)
		assert.Truef(t, !called, "called")
		assert.ErrorIs(t, validate.Error{}, err)
	})

	t.Run("no upgrade validate.Error", func(t *testing.T) {
		// arrange
		var (
			called   = false
			expected = validate.Errorf("already validation error")
			arg      = TestArg{ValidateFunc: func() error {
				return expected
			}}
			sut = validate.Middleware(func(ctx context.Context, arg TestArg) error {
				called = true
				return nil
			})
		)

		// act
		err := sut(t.Context(), arg)

		// assert
		assert.Error(t, err)
		assert.Truef(t, !called, "called")
		assert.Equal(t, expected, err)
	})
}
