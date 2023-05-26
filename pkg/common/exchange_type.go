package common

import "strconv"

//int64类型转换为int型
func Int64ToInt(id64 int64) int {
	strInt64 := strconv.FormatInt(id64, 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}

//string 转 int
func StringToInt(string string) int {
	num, _ := strconv.Atoi(string)
	return num
}

//string 转 int64
func StringToInt64(string string) int64 {
	int64, _ := strconv.ParseInt(string, 10, 64)
	return int64
}

//int64 转 string：
func Int64ToString(int64 int64) string {
	return strconv.FormatInt(int64, 10)
}

//int 转 string
func IntToString(int int) string {
	return strconv.Itoa(int)
}

//string到float32(float64)
func StringToFloat(string string, bitSize int) float64 {
	float, _ := strconv.ParseFloat(string, bitSize)
	return float
}

//float到string  bitSize 32/64
func FloatToString(float float64, bitSize int) string {
	return strconv.FormatFloat(float, 'E', -1, bitSize)
}
