// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package validate_test

import (
	"github.com/kyuff/validate"
	"sync"
)

// Ensure, that ValidatorMock does implement validate.Validator.
// If this is not the case, regenerate this file with moq.
var _ validate.Validator = &ValidatorMock{}

// ValidatorMock is a mock implementation of validate.Validator.
//
//	func TestSomethingThatUsesValidator(t *testing.T) {
//
//		// make and configure a mocked validate.Validator
//		mockedValidator := &ValidatorMock{
//			ValidateFunc: func() error {
//				panic("mock out the Validate method")
//			},
//		}
//
//		// use mockedValidator in code that requires validate.Validator
//		// and then make assertions.
//
//	}
type ValidatorMock struct {
	// ValidateFunc mocks the Validate method.
	ValidateFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Validate holds details about calls to the Validate method.
		Validate []struct {
		}
	}
	lockValidate sync.RWMutex
}

// Validate calls ValidateFunc.
func (mock *ValidatorMock) Validate() error {
	if mock.ValidateFunc == nil {
		panic("ValidatorMock.ValidateFunc: method is nil but Validator.Validate was just called")
	}
	callInfo := struct {
	}{}
	mock.lockValidate.Lock()
	mock.calls.Validate = append(mock.calls.Validate, callInfo)
	mock.lockValidate.Unlock()
	return mock.ValidateFunc()
}

// ValidateCalls gets all the calls that were made to Validate.
// Check the length with:
//
//	len(mockedValidator.ValidateCalls())
func (mock *ValidatorMock) ValidateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockValidate.RLock()
	calls = mock.calls.Validate
	mock.lockValidate.RUnlock()
	return calls
}
