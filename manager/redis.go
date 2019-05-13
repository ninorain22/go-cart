package manager

import (
	"github.com/gomodule/redigo/redis"
	"time"
	"github.com/astaxie/beego"
)

// Redis连接池
var RedisClient *redis.Pool

func init() {
	RedisClient = &redis.Pool{
		MaxIdle: beego.AppConfig.DefaultInt("redis::maxIdle", 20),
		MaxActive: beego.AppConfig.DefaultInt("redis::maxActive", 500),
		IdleTimeout: time.Duration(beego.AppConfig.DefaultInt("redis::idleTimeout", 20)) * time.Second,
		Wait: beego.AppConfig.DefaultBool("redis::wait", true),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.DefaultString("redis::host", ":6379"))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
