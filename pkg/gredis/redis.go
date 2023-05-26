package gredis

import (
	"time"
	"wood/pkg/logging"

	"github.com/gomodule/redigo/redis"

	"wood/pkg/setting"
)

var RedisConn *redis.Pool

// redis
const (
	healthCheckPeriod = time.Second * 30
)

// GetPools 集群所有的redis链接，
// 用途 redis 分布式锁
func GetPools() []*redis.Pool {
	return []*redis.Pool{RedisConn}
}

// Setup Initialize the Redis instance
func Setup() {
	var dialOpts []redis.DialOption
	//if setting.RedisSetting.User != "" {
	//	dialOpts = append(dialOpts, redis.DialUsername(setting.RedisSetting.User))
	//}
	if setting.RedisSetting.Password != "" {
		dialOpts = append(dialOpts, redis.DialPassword(setting.RedisSetting.Password))
	}
	if setting.RedisSetting.DbName > 0 {
		dialOpts = append(dialOpts, redis.DialDatabase(setting.RedisSetting.DbName))
	}

	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		Wait:        setting.RedisSetting.MaxActiveWait,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", setting.RedisSetting.Host, dialOpts...)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	// PING
	if err := Ping(); err != nil {
		logging.Fatal("gredis.Setup err: %v", err)
	}
}

type Redis struct {
}

// Ping .
func Ping() (err error) {
	// PING
	conn := RedisConn.Get()
	defer conn.Close()
	_, err = conn.Do("PING")
	if err != nil {
		return
	}
	return
}
