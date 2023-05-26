package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"wood/pkg/util/convert"
)

// ClientLiveMap type
type ClientLiveMap map[string][]byte

// ClientLiveInfo .
type ClientLiveInfo struct {
	AuthorNumber string `json:"author_number,omitempty"`
	RemoteIP     string `json:"remote_ip,omitempty"`
	AliveTime    int64  `json:"alive_time,omitempty"`
	IsAlive      bool   `json:"is_alive,omitempty"`
}

// ClientComputerInfo .
// 此处的[]byte,必须是proto.Marshal的结果
// 否则在解码时会报错
type ClientComputerInfo map[string][]byte

// ClientCasePlayMap type
type ClientCasePlayMap map[string][]byte

// ClientCasePlay .
type ClientCasePlay struct {
	AuthorNumber string `json:"author_number,omitempty"`
	RemoteIP     string `json:"remote_ip,omitempty"`
	AliveTime    int64  `json:"alive_time,omitempty"`
	CaseUuid     string `json:"case_uuid,omitempty"`
	PlayStatus   int32  `json:"play_status,omitempty"`
	IsPlay       bool   `json:"is_play,omitempty"`
}

// client live
const (
	ClientLiveRedisKey         = "local_backend_client_keepalive"     // redis key
	ClientLiveSecond           = 60                                   // 60秒
	ClientComputerInfoRedisKey = "local_backend_client_computer_info" // redis key
	ClientComputerInfoExpire   = 86400                                // 1天
	ServerLsFilesRedisKey      = "local_backend_srv_ls_files"         // redis key
	ServerLsFilesExpire        = 30                                   // 30s
	ClientCasePlayRedisKey     = "local_backend_client_case_play"     // redis key
	ClientCasePlayLiveSecond   = 60                                   // 60秒
	ClientCasePlayStatusPlay   = 1                                    // case正在回放中
)

// Set a hash key field/value
func HSet(key, field string, data interface{}, expireTime int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if _, err := conn.Do("hset", key, field, value); err != nil {
		return err
	}

	if _, err := conn.Do("expire", key, expireTime); err != nil {
		return err
	}

	return nil
}

// get a hash field
func HGet(key, field string) ([]byte, error) {

	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("hget", key, field))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// HMGet get a hash field
func HMGet(key string, fields ...string) ([][]byte, error) {

	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.ByteSlices(conn.Do("hmget", redis.Args{}.Add(key).AddFlat(fields)...))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// HMGet2 get a hash field
func HMGet2(key string, fields ...string) (interface{}, error) {

	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := conn.Do("hmget", redis.Args{}.Add(key).AddFlat(fields)...)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// set a hash keys values
func HMSet(key string, data interface{}, expireTime int) error {

	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("hmset", redis.Args{}.Add(key).AddFlat(data)...)
	if err != nil {
		return err
	}
	if _, err := conn.Do("expire", key, expireTime); err != nil {
		return err
	}
	return nil
}

// get a hash all
func HGetAll(key string) ([]interface{}, error) {

	conn := RedisConn.Get()
	defer conn.Close()
	v, err := redis.Values(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return v, nil
}

// delete a hash field
func HDelete(key, field string) error {

	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("hdel", key, field)
	if err != nil {
		return err
	}
	return nil
}

// expire a hash
func HExpire(key string) error {

	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("expire", key, 0) // 12小时过期时间
	if err != nil {
		return err
	}
	return nil
}

// hash length
func HLength(key string) int {

	conn := RedisConn.Get()
	defer conn.Close()
	hLen, err := conn.Do("hlen", key)
	if err != nil {
		return 0
	}
	return convert.ToInt(hLen)
}

// hash fields
func HFields(key string) ([]interface{}, error) {

	conn := RedisConn.Get()
	defer conn.Close()
	resKeys, err := redis.Values(conn.Do("hkeys", key))
	if err != nil {
		return nil, err
	}
	return resKeys, nil
}

// hash vals
func HValues(key string) ([]interface{}, error) {

	conn := RedisConn.Get()
	defer conn.Close()
	resValues, err := redis.Values(conn.Do("hvals", key))
	if err != nil {
		return nil, err
	}
	return resValues, nil
}

// hash field exists
func HFieldExists(key, field string) bool {

	conn := RedisConn.Get()
	defer conn.Close()
	isExist, err := conn.Do("hexists", key, field)
	if err != nil {
		return false
	}
	if convert.ToInt(isExist) == 1 {
		return true
	}
	return false
}
