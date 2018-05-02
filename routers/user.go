package routers

import (
	"github.com/gin-gonic/gin"
	"gin-api/controllers/userController"
)

func SetUserRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/users")

	api.POST("/register", userController.Register)

	api.POST("/login", userController.Login)

	api.GET("/captcha/:mobile", userController.GetCaptcha)

	api.GET("/vcode/:mobile", userController.SendSMSCode)

	return r
}
