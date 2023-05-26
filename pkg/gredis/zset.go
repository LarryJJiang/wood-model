package gredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

//有序集合zset数据类型命令
//redis命令：zadd myzset val1 val2 val3 ...
//Zadd(myzset,val1,val2,val3 string,score1,score2,score3 float64)
func Zadd(data ...interface{}) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("zadd", data...)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//redis命令：zrem myzset val1 val2 val3 ...
func Zrem(myzset, val1, val2, val3 string) {
	conn := RedisConn.Get()
	defer conn.Close()
	len, err := redis.Int(conn.Do("zrem", myzset, val1, val2, val3))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("本次有%d个元素被成功删除\n", len)
}

//redis命令：zrange myzset start end [withscores]
func Zrange(myzset string, start, end, flag int) {
	conn := RedisConn.Get()
	defer conn.Close()
	var vals []interface{}
	var err error

	if flag == 0 {
		//不加withscores
		vals, err = redis.Values(conn.Do("zrange", myzset, start, end))
	} else if flag == 1 {
		//加withscores
		vals, err = redis.Values(conn.Do("zrange", myzset, start, end, "withscores"))
	}

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：zrevrange myzset start end [withscores]
func Zrevrange(myzset string, start, end, flag int) {
	conn := RedisConn.Get()
	defer conn.Close()
	var vals []interface{}
	var err error

	if flag == 0 {
		//不加withscores
		vals, err = redis.Values(conn.Do("zrevrange", myzset, start, end))
	} else if flag == 1 {
		//加withscores
		vals, err = redis.Values(conn.Do("zrevrange", myzset, start, end, "withscores"))
	}

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}

}

//redis命令：redis zrangebyscore myzset start end [withscores]
func Zrangebyscore(myzset string, start, end, flag int) {
	conn := RedisConn.Get()
	defer conn.Close()
	var vals []interface{}
	var err error

	if flag == 0 {
		//不加withscores
		vals, err = redis.Values(conn.Do("zrangebyscore", myzset, start, end))
	} else if flag == 1 {
		//加withscores
		vals, err = redis.Values(conn.Do("zrangebyscore", myzset, start, end, "withscores"))
	}

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：zcard myzset
func Zcard(myzset string) {
	conn := RedisConn.Get()
	defer conn.Close()
	len, err := redis.Int(conn.Do("zcard", myzset))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("有序集合%s中有%d个元素\n", myzset, len)
}

//redis命令：zcount myzset minscore maxscore
func Zcount(myzset string, minscore, maxscore float64) {
	conn := RedisConn.Get()
	defer conn.Close()
	len, err := redis.Int(conn.Do("zcount", myzset, minscore, maxscore))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("有序集合%s位于%.2f和%.2f分数区间内的元素有%d个\n", myzset, minscore, maxscore, len)
}

//redis命令：zrank myzset val
func Zrank(myzset, val string) {
	conn := RedisConn.Get()
	defer conn.Close()
	index, err := redis.Int(conn.Do("zrank", myzset, val))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("有序集合%s中的元素%s对应的索引为%d\n", myzset, val, index)
}

//redis命令：zscore myzset val
func Zscore(myzset, val string) {
	conn := RedisConn.Get()
	defer conn.Close()
	score, err := redis.Float64(conn.Do("zscore", myzset, val))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("有序集合%s中的元素%s对应的分数为%.2f\n", myzset, val, score)
}
