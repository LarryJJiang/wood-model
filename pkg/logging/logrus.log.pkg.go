package logging

import (
	golog "github.com/ikaiguang/go-utils/log"
	"github.com/sirupsen/logrus"
	"github.com/wood/models"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
	"wood/pkg/setting"
)

// logHandler .
var logHandler *logrus.Logger

// Logger .
func Logger() *logrus.Logger {
	return logHandler
}

// Setup .
func Setup() {
	var err error
	logHandler, err = NewLogHandler(setting.IsOutputDebug())
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
}

// NewLogHandler .
func NewLogHandler(outputDebug bool) (handler *logrus.Logger, err error) {
	var cfg = &golog.Config{
		MysqlEnable:           false,                // mysql
		FileEnable:            true,                 // file system
		FileRotation:          true,                 // file rotation
		FileName:              "",                   // filename
		FileOptTimeLocation:   time.Local,           // default local (default: rotatelogs.Local)
		FileOptLinkName:       "",                   // link name (default: "")
		FileOptForceNewFile:   false,                // force new file (default: false)
		FileOptMaxAge:         time.Hour * 24 * 365, // lifetime (default: 7 days)
		FileOptRotationTime:   time.Hour * 24,       // rotation time(default: 86400 sec)
		FileOptRotationCount:  0,                    // rotation count (default: -1)
		FileOptRotationSuffix: "_%Y%m%d",            // rotation suffix(example:path+".%Y_%m_%d-%H_%M_%S.log")
	}
	cfg.FileName = filepath.Join(setting.AppSetting.LogSavePath, setting.AppSetting.LogSaveName)
	if !filepath.IsAbs(cfg.FileName) {
		cfg.FileName = filepath.Join(setting.AppSetting.RuntimeRootPath, cfg.FileName)
	}
	cfg.FileOptRotationSuffix += "." + strings.TrimPrefix(setting.AppSetting.LogFileExt, ".")

	// init
	handler, err = golog.NewLogWithConfig(cfg)
	if err != nil {
		return
	}
	handler.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05", ForceColors: true})
	//handler.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	// debug
	if !outputDebug {
		handler.SetOutput(ioutil.Discard)
	}
	return
}

// NewErrLogger .
func NewErrLogger() *ErrLogger {
	return &ErrLogger{logger: logHandler}
}

// ErrLogger .
type ErrLogger struct {
	logger *logrus.Logger
}

// Printf .
func (s *ErrLogger) Printf(format string, v ...interface{}) {
	s.logger.Errorf(format, v...)
}

// NewInfoLogger .
func NewInfoLogger() *InfoLogger {
	return &InfoLogger{logger: logHandler}
}

// InfoLogger .
type InfoLogger struct {
	logger *logrus.Logger
}

// Printf .
func (s *InfoLogger) Printf(format string, v ...interface{}) {
	s.logger.Infof(format, v...)
}

// NewTraceLogger .
func NewTraceLogger() *TraceLogger {
	return &TraceLogger{logger: logHandler}
}

// TraceLogger .
type TraceLogger struct {
	logger *logrus.Logger
}

// Printf .
func (s *TraceLogger) Printf(format string, v ...interface{}) {
	s.logger.Tracef(format, v...)
}
