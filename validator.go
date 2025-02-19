package validate

type Validator interface {
	Validate() error
}

type ValidatorFunc func() error

func (fn ValidatorFunc) Validate() error {
	return fn()
}
