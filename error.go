package validate

import (
	"errors"
	"fmt"
)

type Error struct {
	err error
}

func Errorf(template string, args ...any) error {
	return Error{
		err: fmt.Errorf(template, args...),
	}
}

func (err Error) Error() string {
	if err.err == nil {
		return "<nil>"
	}

	return err.err.Error()
}

func (err Error) Unwrap() error {
	return err.err
}

func (err Error) Is(target error) bool {
	if target == nil {
		return false
	}

	var e Error
	ok := errors.As(target, &e)
	return ok
}
