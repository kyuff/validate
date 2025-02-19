package assert_test

import (
	"errors"
	"testing"

	"github.com/kyuff/validate/internal/assert"
)

func TestAsserts(t *testing.T) {
	var testCases = []struct {
		name   string
		assert func(t *testing.T)
		failed bool
	}{
		{
			name: "Equal success",
			assert: func(t *testing.T) {
				assert.Equal(t, 1, 1)
			},
			failed: false,
		},
		{
			name: "Equal failed",
			assert: func(t *testing.T) {
				assert.Equal(t, 1, 2)
			},
			failed: true,
		},
		{
			name: "Equal success",
			assert: func(t *testing.T) {
				assert.Equalf(t, 1, 1, "hello test")
			},
			failed: false,
		},
		{
			name: "Equalf failed",
			assert: func(t *testing.T) {
				assert.Equalf(t, 1, 2, "hello test")
			},
			failed: true,
		},
		{
			name: "EqualSlice success",
			assert: func(t *testing.T) {
				assert.EqualSlice(t, []int{1, 2, 3}, []int{1, 2, 3})
			},
			failed: false,
		},
		{
			name: "EqualSlice failed size",
			assert: func(t *testing.T) {
				assert.EqualSlice(t, []int{1, 2, 3}, []int{})
			},
			failed: true,
		},
		{
			name: "EqualSlice failed item",
			assert: func(t *testing.T) {
				assert.EqualSlice(t, []int{1, 2, 3}, []int{1, 2, 4})
			},
			failed: true,
		},
		{
			name: "Truef success",
			assert: func(t *testing.T) {
				assert.Truef(t, true, "hello test")
			},
			failed: false,
		},
		{
			name: "Truef failed",
			assert: func(t *testing.T) {
				assert.Truef(t, false, "hello test")
			},
			failed: true,
		},
		{
			name: "NoError success",
			assert: func(t *testing.T) {
				assert.NoError(t, nil)
			},
			failed: false,
		},
		{
			name: "NoError failed",
			assert: func(t *testing.T) {
				assert.NoError(t, errors.New("error"))
			},
			failed: true,
		},
		{
			name: "Error success",
			assert: func(t *testing.T) {
				assert.Error(t, errors.New("error"))
			},
			failed: false,
		},
		{
			name: "Error failed",
			assert: func(t *testing.T) {
				assert.Error(t, nil)
			},
			failed: true,
		},
		{
			name: "Panic success",
			assert: func(t *testing.T) {
				assert.Panic(t, func() {
					panic("test")
				})
			},
			failed: false,
		},
		{
			name: "Panic failed",
			assert: func(t *testing.T) {
				assert.Panic(t, func() {
					// ... no panic
				})
			},
			failed: true,
		},
		{
			name: "NoPanic success",
			assert: func(t *testing.T) {
				assert.NoPanic(t, func() {
					// ... no panic
				})
			},
			failed: false,
		},
		{
			name: "NoPanic failed",
			assert: func(t *testing.T) {
				assert.NoPanic(t, func() {
					panic("test")
				})
			},
			failed: true,
		},

		{
			name: "Match success",
			assert: func(t *testing.T) {
				assert.Match(t, "^[A-Z][0-9].*$", "B5something")
			},
			failed: false,
		},
		{
			name: "Match failed",
			assert: func(t *testing.T) {
				assert.Match(t, "^[A-Z][0-9].*$", "!B5something")
			},
			failed: true,
		},
		{
			name: "Match failed compile",
			assert: func(t *testing.T) {
				assert.Match(t, "(", "RE cannot compile")
			},
			failed: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// arrange
			var (
				x = &testing.T{}
			)

			// act
			testCase.assert(x)

			// arrange
			assert.Equal(t, testCase.failed, x.Failed())
		})
	}
}
