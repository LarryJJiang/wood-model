package logging

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strconv"
)

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logHandler.WithFields(withFields()).Debugln(v...)
}

// Voice .
func Voice(v ...interface{}) {
	logHandler.WithFields(logrus.Fields{"voice_search": "voicevoice_search"}).Infoln(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logHandler.WithFields(withFields()).Infoln(v...)
}

// Infof output logs at info level
func Infof(format string, v ...interface{}) {
	logHandler.WithFields(withFields()).Infof(format, v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logHandler.WithFields(withFields()).Warnln(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logHandler.WithFields(withFields()).Errorln(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logHandler.WithFields(withFields()).Fatalln(v...)
}

// withFields .
func withFields() logrus.Fields {
	return logrus.Fields{"file": file()}
}

// file .
func file() string {
	_, file, line, _ := runtime.Caller(3)

	return filepath.Base(file) + ":" + strconv.Itoa(line)
}
