package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"woods/pkg/logging"
	"woods/pkg/util/convert"
)

func GetQueryToStrE(c *gin.Context, key string) (string, error) {
	str, ok := c.GetQuery(key)
	if !ok {
		return "", errors.New("没有这个值传入")
	}
	return str, nil
}

func GetQueryToStr(c *gin.Context, key string, defaultValues ...string) string {
	str, err := GetQueryToStrE(c, key)
	if str == "" || err != nil {
		var defaultValue string
		if len(defaultValues) > 0 {
			defaultValue = defaultValues[0]
		}
		return defaultValue
	}
	return str
}

func GetQueryToUintE(c *gin.Context, key string) (uint, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUintE(str)
}

func GetQueryToIntE(c *gin.Context, key string) (int, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToIntE(str)
}

func GetQueryToUint(c *gin.Context, key string, defaultValues ...uint) uint {
	var defaultValue uint
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUintE(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

func GetQueryToUint64E(c *gin.Context, key string) (uint64, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToUint64E(str)
}

func GetQueryToInt64E(c *gin.Context, key string) (int64, error) {
	str, err := GetQueryToStrE(c, key)
	if err != nil {
		return 0, err
	}
	return convert.ToInt64(str), nil
}

// GetQueryToUint64 将获取的值转为uint64
func GetQueryToUint64(c *gin.Context, key string, defaultValues ...uint64) uint64 {
	var defaultValue uint64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToUint64E(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetQueryToInt 将获取的值转为int
func GetQueryToInt(c *gin.Context, key string, defaultValues ...int) int {
	val, err := GetQueryToIntE(c, key)
	if err != nil {
		var defaultValue int
		if len(defaultValues) > 0 {
			defaultValue = defaultValues[0]
		}
		return defaultValue
	}
	return val
}

// GetQueryToBool 将获取的值转为bool
func GetQueryToBool(c *gin.Context, key string, defaultValues ...bool) bool {
	str, ok := c.GetQuery(key)
	if !ok || str == "" {
		var defaultValue bool
		if len(defaultValues) > 0 {
			defaultValue = defaultValues[0]
		}
		return defaultValue
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		logging.Logger().Error("参数解析错误，对应的值不是布尔类型：", str)
		return false
	}
	return b
}

// GetQueryToUint64 将获取的值转为uint64
func GetQueryToInt64(c *gin.Context, key string, defaultValues ...int64) int64 {
	var defaultValue int64
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	val, err := GetQueryToInt64E(c, key)
	if err != nil {
		return defaultValue
	}
	return val
}
