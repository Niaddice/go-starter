package redis

import (
	"github.com/go-redis/redis"
	"go-starter/config"
)

// 声明一个全局的rdb变量
var Rdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig("redis.url").(string),
		Password: config.GetConfig("redis.password").(string), // no password set
		DB:       0,                                           // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
