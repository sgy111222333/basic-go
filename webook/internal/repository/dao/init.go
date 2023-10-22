package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	// 严格来说, 这不是优秀实践, 建表应该走流程
	return db.AutoMigrate(&User{}) // 传入dao.User
}
