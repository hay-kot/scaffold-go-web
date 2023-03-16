// Package validate provides a wrapper around the go-playground/validator
package validate

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Check checks a struct for validation errors and returns any errors the occur. This
// wraps the validate.Struct() function and provides some error wrapping. When
// a validator.ValidationErrors is returned, it is wrapped transformed into a
// FieldErrors array and returned.
func Check(val any) error {
	var once sync.Once
	once.Do(func() {
		validate = validator.New()
	})

	err := validate.Struct(val)

	if err != nil {
		verrors, ok := err.(validator.ValidationErrors) // nolint:errorlint
		if !ok {
			return err
		}

		fields := make(FieldErrors, 0, len(verrors))
		for _, verr := range verrors {
			field := FieldError{
				Field: verr.Field(),
				Error: verr.Error(),
			}

			fields = append(fields, field)
		}
		return fields
	}

	return nil
}
