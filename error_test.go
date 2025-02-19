package validate_test

import (
	"context"
	"errors"
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

func TestError(t *testing.T) {

	t.Run("no panic on empty error", func(t *testing.T) {
		// arrange
		var (
			sut = validate.Error{}
		)

		// act
		msg := sut.Error()

		// assert
		assert.Equal(t, msg, "<nil>")
	})

	t.Run("contains message", func(t *testing.T) {
		// arrange
		var (
			sut = validate.Errorf("my message: %d", 5)
		)

		// act
		msg := sut.Error()

		// assert
		assert.Equal(t, msg, "my message: 5")
	})

	t.Run("support unwrap", func(t *testing.T) {
		// arrange
		var (
			sut = validate.Errorf("has context.Canceled: %w", context.Canceled)
		)

		// act
		got := errors.Is(sut, context.Canceled)

		// assert
		assert.Truef(t, got, "should support unwrap")
	})
}
