//go:build !k8s

package config

var Config = config{
	DB: DBConfig{
		DSN: "root:root@tcp(localhost:13306)/webook?charset=utf8mb4",
	},
	Redis: RedisConfig{
		Addr: "127.0.0.1:16379",
	},
}
