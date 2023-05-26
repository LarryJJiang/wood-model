package gredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

//无序集合set数据类型命令
//redis命令：sadd myset val1 val2 val3 ...
func Sadd(myset, val1, val2, val3 string) {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("sadd", myset, val1, val2, val3)

	if err != nil {
		log.Fatal(err)
	}
}

//redis命令：srem myset val
func Srem(myset, val string) {
	conn := RedisConn.Get()
	defer conn.Close()
	//c.Do删除成功返回1，然后用redis.Bool转换成bool值true
	isDel, err := redis.Bool(conn.Do("srem", myset, val))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(isDel)
}

//redis命令：spop myset
func Spop(myset string) {
	conn := RedisConn.Get()
	defer conn.Close()
	val, err := redis.String(conn.Do("spop", myset))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("随机删除一个元素并返回：" + val)
}

//redis命令：smembers myset
func Smembers(myset string) {
	conn := RedisConn.Get()
	defer conn.Close()
	vals, err := redis.Values(conn.Do("smembers", myset))

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：scard myset
func Scard(myset string) {
	conn := RedisConn.Get()
	defer conn.Close()
	len, err := redis.Int(conn.Do("scard", myset))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("集合长度为：%d\n", len)
}

//redis命令：sismember myset val
func Sismember(myset, val string) {
	conn := RedisConn.Get()
	defer conn.Close()
	isMember, err := redis.Bool(conn.Do("sismember", myset, val))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(isMember)
}

//redis命令：srandmember myset count
func Srandmember(myset string, count int) {
	conn := RedisConn.Get()
	defer conn.Close()
	vals, err := redis.Values(conn.Do("srandmember", myset, count))

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：smove myset myset2 val
func Smove(myset, myset2, val string) {
	conn := RedisConn.Get()
	defer conn.Close()

	isMoveSuccessful, err := redis.Bool(conn.Do("smove", myset, myset2, val))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("从集合%s中移动元素%s到集合%s，移动结果："+strconv.FormatBool(isMoveSuccessful)+"\n", myset, val, myset2)
}

//redis命令：sunion myset myset2 ...
func Sunion(myset, myset2 string) {
	conn := RedisConn.Get()
	defer conn.Close()

	vals, err := redis.Values(conn.Do("sunion", myset, myset2))

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sunionstore myset3 myset myset2 ...
func Sunionstore(myset, myset2, myset3 string) {
	conn := RedisConn.Get()
	defer conn.Close()

	//将集合myset、myset2取并集后存入myset3集合
	len, err := redis.Int(conn.Do("sunionstore", myset3, myset, myset2))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("集合%s与集合%s取交集后得到的集合%s有%d个元素\n", myset, myset2, myset3, len)

	//打印集合myset3验证结果
	Smembers("myset3")
}

//redis命令：sinter myset myset2 ...
func Sinter(myset, myset2 string) {
	conn := RedisConn.Get()
	defer conn.Close()

	vals, err := redis.Values(conn.Do("sinter", myset, myset2))

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sinterstore myset3 myset myset2 ...
func Sinterstore(myset, myset2, myset3 string) {
	conn := RedisConn.Get()
	defer conn.Close()

	len, err := redis.Int(conn.Do("sinterstore", myset3, myset, myset2))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("集合%s与集合%s取交集后得到的集合%s有%d个元素\n", myset, myset2, myset3, len)

	//打印集合myset3验证结果
	Smembers("myset3")
}

//redis命令：sdiff myset myset2 ...
func Sdiff(myset, myset2 string) {
	conn := RedisConn.Get()
	defer conn.Close()

	vals, err := redis.Values(conn.Do("sdiff", myset, myset2))

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range vals {
		fmt.Printf("k = %v v = %s\n", k, v)
	}
}

//redis命令：sdiffstore myset3 myset myset2 ...
func Sdiffstore(myset, myset2, myset3 string) {
	conn := RedisConn.Get()
	defer conn.Close()

	len, err := redis.Int(conn.Do("sdiffstore", myset3, myset, myset2))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("集合%s与集合%s取交集后得到的集合%s有%d个元素\n", myset, myset2, myset3, len)

	//打印集合myset3验证结果
	Smembers("myset3")

}
