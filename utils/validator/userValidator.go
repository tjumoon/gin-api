package validator

import (
	"gopkg.in/go-playground/validator.v8"
	"regexp"
	"reflect"
)

const MobileRegex  = "^1[0-9]{10}$"
const PasswordRegex = "^[0-9a-zA-Z@.]{6,30}$"

func mobileValidator(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	mobile := field.String()
	return validatorMobile(mobile)
}

func validatorMobile(mobile string) bool {
	b, _ := regexp.MatchString(MobileRegex, mobile)
	return b
}

func passwordValidator(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	password := field.String()
	return validatorPassword(password)
}

func validatorPassword(password string) bool {
	b, _ := regexp.MatchString(PasswordRegex, password)
	return b
}