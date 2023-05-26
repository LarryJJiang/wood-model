package gredis

import (
	"github.com/gomodule/redigo/redis"
)

// 从列表左侧插入值
func LPush(key, value string, expireTime int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	if _, err := conn.Do("lpush", key, value); err != nil {
		return err
	}
	if _, err := conn.Do("expire", key, expireTime); err != nil {
		return err
	}
	return nil
}

// 从列表右侧插入值
func RPush(key, value string, expireTime int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	if _, err := conn.Do("rpush", key, value); err != nil {
		return err
	}
	if _, err := conn.Do("expire", key, expireTime); err != nil {
		return err
	}
	return nil
}

// 设置列表所有的值
func LMpush(key string, values interface{}, expireTime int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("lpush", redis.Args{}.Add(values)...)
	if err != nil {
		return err
	}
	if _, err := conn.Do("expire", key, expireTime); err != nil {
		return err
	}
	return nil
}

// 获取列表所有值
func LGetAll(key string) ([]interface{}, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	num, err := conn.Do("llen", key)
	if err != nil {
		return nil, err
	}

	values, err := redis.Values(conn.Do("lrange", key, 0, num))
	if err != nil {
		return nil, err
	}
	return values, nil
}

// 删除指定值
func LRem(key, value string) error {

	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("lrem", key, 0, value)
	if err != nil {
		return err
	} else {
		return nil
	}
}
