package logger

import (
	"fmt"
	"os"
	"time"
)

const (
	warn  = "warn"
	info  = "info"
	debug = "debug"
	error = "error"
)

// Debug デバッグ用
func Debug(v interface{}) {
	os.Stdout.Write([]byte(logText(debug, v)))
}

// Info サーバの情報
func Info(v interface{}) {
	os.Stdout.Write([]byte(logText(info, v)))
}

// Warn 警告
func Warn(v interface{}) {
	os.Stdout.Write([]byte(logText(warn, v)))
}

// Error サーバのエラー
func Error(v interface{}) {
	os.Stdout.Write([]byte(logText(error, v)))
}

func logText(logLevel string, v interface{}) string {
	return fmt.Sprintf("%s [%s] %v\n", time.Now().Format("2006/01/02 03:04:05"), logLevel, v)
}
