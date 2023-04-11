package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-fiber-postgres/config"
	"go-fiber-postgres/exception"
	"go-fiber-postgres/model"
	"regexp"
)

func Validate[T any](val T) (bool, error) {
	var errors validator.ValidationErrors
	validate := config.NewValidator()
	err := validate.Struct(val)
	if err != nil {
		var errorx exception.ValidationError
		for _, errors := range err.(validator.ValidationErrors) {
			if errorx.Message == "" {
				errorx = exception.ValidationError{
					Message: fmt.Sprintf("{\"%s\": \"%v\"}", errors.Field(), FormatValidationError(errors.Tag(), errors.Param(), errors.Value())),
				}
			} else {
				errorx = exception.ValidationError{
					Message: errorx.Message[:len(errorx.Message)-1] + "," + fmt.Sprintf("\"%s\": \"%v\"}", errors.Field(), FormatValidationError(errors.Tag(), errors.Param(), errors.Value())),
				}
			}

		}
		return true, errorx
	}
	return false, errors
}

func FormatValidationError(tag string, param interface{}, value interface{}) string {
	switch tag {
	case "required":
		return model.NOT_BLANK_ERR_TYPE
	case "alphanum":
		return model.MUST_STRING_ERR_TYPE
	case "number":
		return model.MUST_NUMBER_ERR_TYPE
	case "status-validation":
		return model.NOT_VALID_ERR_TYPE
	case "is-negative":
		return model.CANNOT_BE_NEGATIVE_NUMBER
	case "lat-long":
		return model.NOT_VALID_ERR_TYPE
	case "gte":
		return model.Min(param)
	case "lte":
		return model.Max(param)
	case "special-gte":
		return model.Min(param)
	case "special-lte":
		return model.Max(param)
	case "date":
		return model.NOT_VALID_ERR_TYPE
	case "date-format":
		return model.NOT_VALID_ERR_TYPE
	case "special-char":
		return model.NOT_VALID_ERR_TYPE
	case "phone-number":
		return model.NOT_VALID_ERR_TYPE
	case "escape-unicode":
		var regex, _ = regexp.Compile(`[\\|'"?^*%]`)
		if regex.FindString(value.(string)) != "" {
			return model.NOT_VALID_ERR_TYPE
		}
		return ""
	case "not-minus":
		return model.NOT_VALID_ERR_TYPE
	case "address":
		return model.NOT_VALID_ERR_TYPE
	default:
		return tag
	}
}
