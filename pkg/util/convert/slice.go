package convert

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// []int转成字符串
func Join2str(slice []int, sep string) string {
	s := fmt.Sprintf("%v", slice)
	sliceString := regexp.MustCompile(`[\w.]+`).FindAllString(s, -1)
	return strings.Join(sliceString, sep)
}

// 获取不重复的切片
func UnrepeatedSlice(slice []int) []int {
	length := len(slice)
	sliceMap := make(map[int]int, 0)
	newSlice := make([]int, 0)
	for i := 0; i < length; i++ {
		_, ok := sliceMap[slice[i]]
		if ok {
			continue
		}
		newSlice = append(newSlice, slice[i])
		sliceMap[slice[i]] = 0
	}
	return newSlice
}

func GetFieldSlice(m interface{}) []string {
	v := reflect.TypeOf(m)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	data := make([]string, v.NumField())
	data = []string{"id", "create_time", "update_time"}
	for i := 0; i < v.NumField(); i++ {
		var field = v.Field(i).Tag.Get("gorm")
		if field != "" {
			fieldArray := strings.Split(field, ";")
			for _, value := range fieldArray {
				if strings.Contains(value, "column:") {
					field = strings.Replace(value, "column:", "", 1)
					data = append(data, fmt.Sprintf("`%v`", field))
				}
			}
		}
	}
	return data
}
