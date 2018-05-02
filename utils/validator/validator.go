package validator

import (
	"github.com/gin-gonic/gin/binding"
)

func InitValidator()  {
	binding.Validator.RegisterValidation("password", passwordValidator)
	binding.Validator.RegisterValidation("mobile", mobileValidator)
}



