package ioc

import (
	"github.com/sgy111222333/basic-go/webook/config"
	"github.com/sgy111222333/basic-go/webook/internal/repository/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook?charset=utf8mb4"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug() // 打印出所有生成的sql语句
	err = dao.InitTables(db)
	if err != nil {
		return nil
	}
	return db
}
