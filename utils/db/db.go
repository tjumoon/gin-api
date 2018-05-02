package db

import (
	"gin-api/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"gin-api/utils/log"
)

var ORM *gorm.DB

func InitDB()  {
	var (
		err error
		dbType, dbName, user, password, host string
	)
	sec, err := config.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	//tablePrefix = sec.Key("TABLE_PREFIX").String()

	ORM, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		log.Fatalf("DB init fail %v", err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}
	ORM.SingularTable(true)
	ORM.DB().SetMaxIdleConns(10)
	ORM.DB().SetMaxOpenConns(100)
	if config.RunMode == "debug" {
		ORM.LogMode(true)
	} else {
		ORM.LogMode(false)
	}

}

func CloseDB()  {
	defer ORM.Close()
}
