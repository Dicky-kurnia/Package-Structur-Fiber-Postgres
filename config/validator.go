package config

import (
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	//validate           *validator.Validate
	escapeUnicodeRegex = regexp.MustCompile(`[\\|'"?^*%]`)
	latLongRegexString = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?),\\s*[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	latLongRegex       = regexp.MustCompile(latLongRegexString)
	specialCharsRegex  = regexp.MustCompile(`[\` + "`" + `\~\!\@\#\$\&\(\)\-\=\_\+\[\]\{\}\|\;\:\,\.\<\>\/\r\n]`)
	phoneNumberRegex   = regexp.MustCompile(`^(0)8[1-9][0-9]{2,12}$`)
	addressRegex       = regexp.MustCompile(`^[a-zA-Z0-9 \` + "`" + `\~\!\@\#\$\&\(\)\-\=\_\+\[\]\{\}\|\;\:\,\.\<\>\/\r\n]+$`)
)

func NewValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.RegisterValidation("status-validation", statusValidation)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("is-negative", isNegative)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("escape-unicode", escapeUnicode)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("lat-long", latLong)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("address", address)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("date-format", validateDate)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("special-char", isSpecialChar)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("special-gte", specialGte)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("special-lte", specialLte)
	if err != nil {
		panic(err)
	}

	err = validate.RegisterValidation("phone-number", phoneNumber)
	if err != nil {
		panic(err)
	}
	err = validate.RegisterValidation("not-minus", notMinus)
	if err != nil {
		panic(err)
	}

	return validate
}
func latLong(fl validator.FieldLevel) bool {
	return latLongRegex.MatchString(fl.Field().String())
}

func isNegative(fl validator.FieldLevel) bool {
	return !math.Signbit(float64(fl.Field().Int()))
}

func escapeUnicode(fl validator.FieldLevel) bool {
	return escapeUnicodeRegex.FindString(fl.Field().String()) == ""
}

func isSpecialChar(fl validator.FieldLevel) bool {
	return !specialCharsRegex.MatchString(fl.Field().String())
}

func address(fl validator.FieldLevel) bool {
	return addressRegex.MatchString(fl.Field().String())
}

func validateDate(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}

func specialGte(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	length := len([]rune(fl.Field().String()))
	param, _ := strconv.Atoi(fl.Param())
	if length >= param {
		return true
	}
	return false
}

func specialLte(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return true
	}
	length := len([]rune(fl.Field().String()))
	param, _ := strconv.Atoi(fl.Param())
	if length <= param {
		return true
	}
	return false
}

func phoneNumber(fl validator.FieldLevel) bool {
	return phoneNumberRegex.MatchString(fl.Field().String())
}

func notMinus(fl validator.FieldLevel) bool {
	return !math.Signbit(fl.Field().Float())
}

func statusValidation(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case "SUCCESS":
		return true
	case "ONPROCESS":
		return true
	case "PAID":
		return true
	case "CANCELLED":
		return true
	case "SENDING":
		return true
	case "DELIVERED":
		return true
	default:
		return false
	}
}
