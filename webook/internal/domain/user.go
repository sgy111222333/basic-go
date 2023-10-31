package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string // 也可以用[]byte, 不过string显示更友好
	Nickname string
	Birthday time.Time
	AboutMe  string
	// 一律用UTC+0的时区, 只在返回给前端展示的时候转成对应时区
	Ctime time.Time
}

//type Address struct {
//	Province string
//	Region   string
//}
