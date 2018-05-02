package common

import (
	"gin-api/utils/redis"
	"gin-api/utils/log"
	"gin-api/utils/db"
)

func StartUp()  {
	log.Init()
	redis.InitRedisPool()
	db.InitDB()
}


