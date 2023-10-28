//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		// * 因为是k8s内部互连, 用 名字:端口 就可以, 这个端口不是targetPort
		DSN: "root:root@tcp(webook-mysql-service:3308)/webook?charset=utf8mb4",
	},
	Redis: RedisConfig{
		Addr: "webook-redis-service:6380",
	},
}
