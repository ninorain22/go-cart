package manager

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// Redis连接池
var RedisClient *redis.Pool

func init() {
	// todo: 建议从配置文件app.conf中获取配置参数
	RedisClient = &redis.Pool{
		MaxIdle: 10,
		MaxActive: 500,
		IdleTimeout: 20 * time.Second,
		Wait: true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
