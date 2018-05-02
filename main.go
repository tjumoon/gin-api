package main

import (
	"github.com/fvbock/endless"
	"fmt"
	"syscall"
	"gin-api/routers"
	"gin-api/utils/config"
	"gin-api/common"
	"gin-api/utils/log"
)


// @title Gin API
// @version 1.0
// @description This is a gin-api server Petstore server.

// @contact.name API Support
// @contact.url http://www.simonblog.cn
// @contact.email simon_yang@aliyun.com

func main()  {

	common.StartUp()
	endless.DefaultReadTimeOut = config.ReadTimeout
	endless.DefaultWriteTimeOut = config.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", config.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Infof("Actual pid is %d\n", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Errorf("Server start error: %v", err)
	}
}
