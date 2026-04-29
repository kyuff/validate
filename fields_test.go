package validate_test

import (
	"testing"

	"github.com/kyuff/validate"
	"github.com/kyuff/validate/internal/assert"
)

type fieldName string

func (n fieldName) Validate() error {
	return validate.StringNotEmpty(n).Validate()
}

type fieldAddress string

func (a fieldAddress) Validate() error {
	return validate.StringNotEmpty(a).Validate()
}

type fieldOrder struct {
	Name    fieldName
	Address fieldAddress
}

type fieldOrderWithInvalidField struct {
	Name  fieldName
	Score int
}

type fieldLargeOrder struct {
	Name        fieldName
	Address     fieldAddress
	City        fieldName
	Country     fieldName
	PostalCode  fieldName
	PhoneNumber fieldName
	Email       fieldName
	Notes       fieldAddress
}

func BenchmarkFieldsStrict(b *testing.B) {
	b.Run("small struct", func(b *testing.B) {
		sut := fieldOrder{
			Name:    "John",
			Address: "123 Main St",
		}

		b.ResetTimer()
		for b.Loop() {
			_ = validate.FieldsStrict(sut).Validate()
		}
	})

	b.Run("large struct", func(b *testing.B) {
		sut := fieldLargeOrder{
			Name:        "John",
			Address:     "123 Main St",
			City:        "Springfield",
			Country:     "US",
			PostalCode:  "12345",
			PhoneNumber: "555-1234",
			Email:       "john@example.com",
			Notes:       "Leave at door",
		}

		b.ResetTimer()
		for b.Loop() {
			_ = validate.FieldsStrict(sut).Validate()
		}
	})

	b.Run("pointer to struct", func(b *testing.B) {
		sut := &fieldOrder{
			Name:    "John",
			Address: "123 Main St",
		}

		b.ResetTimer()
		for b.Loop() {
			_ = validate.FieldsStrict(sut).Validate()
		}
	})
}

func TestFieldsStrict(t *testing.T) {
	t.Run("fail when a field does not implement Validator", func(t *testing.T) {
		// arrange
		var (
			sut = fieldOrderWithInvalidField{
				Name:  "John",
				Score: 10,
			}
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.Error(t, got)
		assert.ErrorIs(t, validate.Error{}, got)
	})

	t.Run("fail when a field is invalid", func(t *testing.T) {
		// arrange
		var (
			sut = fieldOrder{
				Name:    "",
				Address: "123 Main St",
			}
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.Error(t, got)
	})

	t.Run("fail when multiple fields are invalid", func(t *testing.T) {
		// arrange
		var (
			sut = fieldOrder{
				Name:    "",
				Address: "",
			}
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.Error(t, got)
	})

	t.Run("fail on nil pointer", func(t *testing.T) {
		// arrange
		var (
			sut *fieldOrder = nil
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.Error(t, got)
		assert.ErrorIs(t, validate.Error{}, got)
	})

	t.Run("fail on non-struct input", func(t *testing.T) {
		// arrange
		var (
			sut = "not a struct"
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.Error(t, got)
		assert.ErrorIs(t, validate.Error{}, got)
	})

	t.Run("pass when all fields are valid", func(t *testing.T) {
		// arrange
		var (
			sut = fieldOrder{
				Name:    "John",
				Address: "123 Main St",
			}
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.NoError(t, got)
	})

	t.Run("pass when given a pointer to a valid struct", func(t *testing.T) {
		// arrange
		var (
			sut = &fieldOrder{
				Name:    "John",
				Address: "123 Main St",
			}
		)

		// act
		got := validate.FieldsStrict(sut).Validate()

		// assert
		assert.NoError(t, got)
	})
}
