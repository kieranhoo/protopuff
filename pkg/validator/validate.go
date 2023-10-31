package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func _Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func Validator(data interface{}) string {
	if errs := _Validate(data); len(errs) > 0 && errs[0].Error {
		errMsg := make([]string, 0)

		for _, err := range errs {
			errMsg = append(errMsg, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return strings.Join(errMsg, " and ")
	}
	return ""
}
