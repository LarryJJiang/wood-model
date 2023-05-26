package gredis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"woods/pkg/logging"
)

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}

func Incr(key string, time int) (times int64, err error) {
	conn := RedisConn.Get()
	defer conn.Close()
	times, err = redis.Int64(conn.Do("INCR", key))
	if err != nil {
		logging.Error("INCR failed:", err)
		return
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			logging.Error("INCR failed:", err.Error())
			return
		}
	}
	return
}

// ExistsAndExpire check a key
func ExistsAndExpire(key string) (exists bool, err error) {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err = redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return
	}
	if !exists {
		return
	}
	_, err = conn.Do("expire", key, 0)

	return
}

func IncrBy(key string, step int, time int) (times int64, err error) {
	conn := RedisConn.Get()
	defer conn.Close()
	times, err = redis.Int64(conn.Do("INCRBY", key, step))
	if err != nil {
		logging.Error("INCR failed:", err)
		return
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			logging.Error("INCR failed:", err.Error())
			return
		}
	}
	return
}

//redis命令：mset key1 val1 key2 val2 key3 val3 ...
func Mset(key ...interface{}) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	fmt.Printf("参数Mset：%v 类型：%T\n", key, key)
	_, err := conn.Do("mset", key...)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//redis命令：mget key1 key2 key3 ...
func Mget(key ...interface{}) []string {
	conn := RedisConn.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("mget", key...))

	if err != nil {
		log.Fatal(err)
	}
	var data []string
	for _, v := range result {
		//fmt.Printf("k = %v  v = %s 类型：%T\n",k, v, v)
		data = append(data, fmt.Sprintf("%s", v))
	}
	return data
}

//redis命令：del key1 key2 ...
func Del(key ...interface{}) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("del", key...)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//redis命令：getrange key start end
func Getrange(key string, start, end int) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("getrange", key, start, end))

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	fmt.Println(s)
	return s, err
}

//redis命令：setex key 10 value
func Setex(key, val string, expire int) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	if _, err := conn.Do("setex", key, expire, val); err != nil {
		log.Fatal("setex error：", err)
		return false
	}
	return true
}

//redis命令：setnx key val
func Setnx(key, val string) interface{} {
	conn := RedisConn.Get()
	defer conn.Close()
	res, err := redis.Bool(conn.Do("setnx", key, val))
	if err != nil {
		log.Fatal("setnx error：", err)
		return false
	}
	fmt.Printf("setnx结果：%v 错误信息：%v  类型：%T", res, err, res)
	return res
}
