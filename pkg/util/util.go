package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
	"wood/pkg/logging"
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}

func GetUUID() string {
	uuidBytes := uuid.NewV4()
	return uuidBytes.String()
}

// 正则验证手机号
func CheckPhone(phone string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

// 生成随机字串
func GetRandomString(sourceStr string, l int) string {

	bytes := []byte(sourceStr)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// File file line
func File(skips ...int) (file string, line int) {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}
	if skip < 0 {
		skip = 0
	}
	_, file, line, _ = runtime.Caller(skip)
	return
}

// BytesCombine : 拼接byte
func BytesCombine(pBytes ...[]byte) []byte {
	arrLen := len(pBytes)
	s := make([][]byte, arrLen)
	for index := 0; index < arrLen; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

// BytesToUint8
func BytesToUint8(b []byte) uint8 {
	_ = b[0]
	return uint8(b[0])
}

// BytesToUint16
func BytesToUint16(b []byte) uint16 {
	_ = b[1]
	return uint16(b[0]) | uint16(b[1])<<8
}

// BytesToUint32
func BytesToUint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// BytesToUint64
func BytesToUint64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
}

// struct 转 map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}
	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			if len(t.Field(i).Tag.Get("json")) > 0 { // 获取public field
				tempField := strings.Split(t.Field(i).Tag.Get("json"), ",")
				data[tempField[0]] = v.Field(i).Interface()
			}
		}
	}
	return data
}

// struct 转 map
func Struct2UrlValues(obj interface{}) url.Values {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = url.Values{}
	for i := 0; i < t.NumField(); i++ {
		data.Set(t.Field(i).Tag.Get("json"), convert.ToString(v.Field(i).Interface()))
		//data[t.Field(i).Tag.Get("json")] = convert.ToString(v.Field(i).Interface())
	}
	return data
}

func Map2List(obj map[string]interface{}) []string {
	keys := make([]string, len(obj))
	var out []string
	i := 0
	for k := range obj {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k != "signature" && k != "ukey_serial_number" {
			out = append(out, fmt.Sprintf("%v=%v", k, obj[k]))
		}
	}
	return out
}

func ReverseBytes(source []byte) []byte {
	byteLen := len(source)
	array := make([]byte, byteLen)
	for i, v := range source {
		array[i] = v
	}
	for i := 0; i < byteLen/2; i++ {
		array[i], array[byteLen-i-1] = array[byteLen-i-1], array[i]
	}
	return array
}

func GenerateDate() string {
	return time.Now().Format("2006-01-02")
}

func ContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// slice 去除重复字符串
func RemoveRepeatedElement(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func LocalIp() []string {
	result := make([]string, 0)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logging.Error(err)
		return result
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	return result
}

//UTF82GBK : transform UTF8 rune into GBK byte array
func UTF82GBK(src string) ([]byte, error) {
	GB18030 := simplifiedchinese.All[0]
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
}

//GBK2UTF8 : transform  GBK byte array into UTF8 string
func GBK2UTF8(src []byte) (string, error) {
	GB18030 := simplifiedchinese.All[0]
	bytes, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	return string(bytes), err
}

// 指定类型数据转化为另一种类型的数据
// 用法 Interface2Interface(v, &out)，不用返回，这个out就是返回
func Interface2Interface(v interface{}, out interface{}) error {
	jsonString, err := json.Marshal(v)
	if err != nil {
		fmt.Println("exchange to json fail")
		return err
	}
	err = json.Unmarshal(jsonString, out)
	if err != nil {
		fmt.Println("exchange to json fail")
		return err
	}
	return nil
}
