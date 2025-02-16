package log

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

var _ Logger = (*consoleLogger)(nil)

// consoleLogger 命令行日志实现
type consoleLogger struct{}

// Debug 日志
func (consoleLogger) Debug(v ...interface{}) {
	output("Debug", fmt.Sprint(v...))
}

// Info 日志
func (consoleLogger) Info(v ...interface{}) {
	output("Info", fmt.Sprint(v...))
}

// Warn 日志
func (consoleLogger) Warn(v ...interface{}) {
	output("Warning", fmt.Sprint(v...))
}

// Error
func (consoleLogger) Error(v ...interface{}) {
	output("Error", fmt.Sprint(v...))
}

// Debugf Debug Format 日志
func (consoleLogger) Debugf(format string, v ...interface{}) {
	output("Debug", fmt.Sprintf(format, v...))
}

// Infof Info Format 日志
func (consoleLogger) Infof(format string, v ...interface{}) {
	output("Info", fmt.Sprintf(format, v...))
}

// Warnf Warning Format 日志
func (consoleLogger) Warnf(format string, v ...interface{}) {
	output("Warning", fmt.Sprintf(format, v...))
}

// Errorf Error Format 日志
func (consoleLogger) Errorf(format string, v ...interface{}) {
	output("Error", fmt.Sprintf(format, v...))
}

func output(level string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	files := strings.Split(file, "/")
	file = files[len(files)-1]

	logFormat := "[%s] %s %s:%d " + fmt.Sprint(v...) + "\n"
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf(logFormat, level, date, file, line)
}
