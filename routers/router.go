package routers

import (
	"github.com/gin-gonic/gin"
	"gin-api/utils/config"
	"gin-api/middleware/log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "gin-api/docs"

	"gin-api/utils/validator"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	switch config.RunMode {
	case "debug":
		r.Use(gin.Logger())
	case "release":
		r.Use(log.AccessLogger())
	}

	validator.InitValidator()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(config.RunMode)

	r.Use(gin.Recovery())

	SetUserRoutes(r)

	return r

}
