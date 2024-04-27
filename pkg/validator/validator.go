package validator

import (
	"errors"
	"strings"

	"github.com/BrondoL/wedding-be/internal/util"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "can not be empty!"
	case "max":
		return "should be less than or equal to " + util.ToSnakeCase(fe.Param()) + "!"
	case "min":
		return "should be greater than or equal to " + util.ToSnakeCase(fe.Param()) + "!"
	case "gte":
		return "should be greater than or equal to " + util.ToSnakeCase(fe.Param()) + "!"
	case "gt":
		return "should be greater than " + util.ToSnakeCase(fe.Param()) + "!"
	case "lte":
		return "should be less than or equal to " + util.ToSnakeCase(fe.Param()) + "!"
	case "email":
		return "must be a valid email address!"
	case "eqfield":
		return "does not match with " + util.ToSnakeCase(fe.Param()) + "!"
	case "ltfield":
		return "must be less than " + util.ToSnakeCase(fe.Param()) + " field!"
	case "gtfield":
		return "must be greater than " + util.ToSnakeCase(fe.Param()) + " field!"
	case "alpha":
		return "must be entirely alphabetic characters!"
	case "alphanum":
		return "must be entirely alpha-numeric characters!"
	case "numeric":
		return "must be an integer!"
	case "oneof":
		return "must be one of " + strings.Replace(fe.Param(), " ", ", ", -1)
	case "len":
		return "must have a length of " + util.ToSnakeCase(fe.Param()) + "!"
	}
	return "something is wrong with this field!"
}

func FormatValidation(err error) interface{} {
	var result []ErrorMsg

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			result = append(result, ErrorMsg{Field: util.ToSnakeCase(fe.Field()), Message: getErrorMsg(fe)})
		}
	}
	if len(result) == 0 {
		return nil
	}

	return result
}
