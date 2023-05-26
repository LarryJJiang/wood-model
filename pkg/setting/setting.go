package setting

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	RuntimeRootPath string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExt   []string
	FileSavePath    string
	FileMaxSize     int
	FileAllowExt    []string

	ExportSavePath string
	FontSavePath   string

	LogSavePath        string
	LogSaveName        string
	LogFileExt         string
	TimeFormat         string
	ConfPath           string
	JwtSecret          string
	JwtIssuer          string
	JwtTokenExpireTime int
	PageSize           int
	TimeZone           string
	Domain             string
	CrontabEnable      bool
	CrontabSpec        string
}

var AppSetting = &App{}

type Server struct {
	OutputDebug  bool
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Host        string
	Port        int
	Name        string
	Password    string
	TablePrefix string
	LogMode     bool
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host          string
	User          string
	Password      string
	DbName        int
	MaxIdle       int
	MaxActive     int
	MaxActiveWait bool
	IdleTimeout   time.Duration
}

var RedisSetting = &Redis{}

type GoogleCloud struct {
	ProjectId   string
	Location    string
	ProcessorId string
	FileUrl     string
}

var GoogleCloudSetting = &GoogleCloud{}

var cfg *ini.File

const (
	CharStr   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberStr = "123456789"
)

var RunMode string

func InitSetUp() {
	var err error
	// 设置环境类型方式
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	RunMode = cfg.Section("").Key("RunMode").String()
	gin.SetMode(RunMode) // 设置模式，有三种模式：debug release test
	fmt.Println("已开启模式：", gin.Mode())
	cfg, err = ini.Load(fmt.Sprintf("conf/app_%s.ini", gin.Mode()))
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app_%s.ini': %v", gin.Mode(), err)
	}
}

// Setup initialize the configuration instance
func SetupConfig() {
	InitSetUp()

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("cloud", GoogleCloudSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	AppSetting.FileMaxSize = AppSetting.FileMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

// IsOutputDebug debug
func IsOutputDebug() bool {
	return ServerSetting.OutputDebug
}
