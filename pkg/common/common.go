package common

import (
	"github.com/gin-gonic/gin"
	"os"
)

// GetPageIndex 获取页码
func GetPageIndex(c *gin.Context) uint64 {
	return GetQueryToUint64(c, "page", 1)
}

// GetPageSize 获取每页记录数
func GetPageSize(c *gin.Context) uint64 {
	pageSize := GetQueryToUint64(c, "page_size", 20)
	if pageSize > 500 {
		pageSize = 20
	}
	return pageSize
}

// GetSort 获取排序信息
func GetSort(c *gin.Context) string {
	return GetQueryToStr(c, "sort")
}

// GetKeyword 获取搜索关键词信息
func GetKeyword(c *gin.Context) string {
	return GetQueryToStr(c, "keyword")
}

// PathExists 路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
