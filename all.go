package validate

import "errors"

func All(validators ...Validator) error {
	var err error
	for _, v := range validators {
		if v == nil {
			continue
		}

		err = errors.Join(err, v.Validate())
	}

	return err
}
