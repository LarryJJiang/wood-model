package convert

import (
	"encoding/binary"
	"reflect"
)

// int64 转 byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// byte 转 int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func StructToMap(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)
	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}
	var data = make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Name] = objValue.Field(i).Interface()
	}
	return data
}
