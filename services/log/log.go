package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"stash-scrapers/common/config"
	"stash-scrapers/common/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type file struct {
	level    map[string]bool
	fileName string
	date     int
	fileFd   *os.File
	err      *log.Logger
	warn     *log.Logger
	info     *log.Logger
	debug    *log.Logger
	info2    *log.Logger
}

var (
	logFile  file
	logF     *os.File
	logTitle = fmt.Sprintf("[%s] ", config.ServiceName)
	logPath  = "./logs"
)

func init() {
	logFile.level = map[string]bool{"Info": true, "Warn": true, "Debug": true, "Error": true}
	logFile.CreateLogFile()
}

// Logs 日志分割
func Logs() gin.HandlerFunc {
	return func(c *gin.Context) {
		if logFile.date != time.Now().Day() {
			logFile.date = time.Now().Day()
			logFile.CreateLogFile()
		}
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		// c.Set(utils.TotalTakeTimeStartKey, start)
		// 业务处理
		c.Next()
		// Stop timer
		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		if raw != "" {
			path = utils.JoinString(path, "?", raw)
		}
		if statusCode >= 200 && statusCode <= 400 {
			Infof("| %3d | %15v | %15s | %-7s %s", statusCode, clientIP, latency, method, path)
		} else if statusCode >= 400 && statusCode < 500 {
			Warnf("| %3d | %15v | %15s | %-7s %s", statusCode, clientIP, latency, method, path)
		} else {
			Errorf("| %3d | %15v | %15s | %-7s %s", statusCode, clientIP, latency, method, path)
		}
	}
}

// Println 输出日常日志
func Println(args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("%s%s%d", filePath, ":", line)
	a := []interface{}{str}
	a = append(a, args...)
	logFile.info2.Println(a...)
}

// Fatalln 输出异常日志
func Fatalln(args ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("%s%s%d", filePath, ":", line)
	a := []interface{}{str}
	a = append(a, args...)
	logFile.info2.Fatalln(a...)
}

// Printf 输出日常日志
func Printf(format string, v ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("%s%s%d", filePath, ":", line)
	// str2 := fmt.Sprintf(format, v...)
	a := []interface{}{str}
	a = append(a, v...)
	logFile.info2.Printf(str, a...)
}

// Printf 输出日常日志
func Fatalf(format string, v ...interface{}) {
	_, filePath, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("%s%s%d", filePath, ":", line)
	// str2 := fmt.Sprintf(format, v...)
	a := []interface{}{str}
	a = append(a, v...)
	logFile.info2.Fatalf(str, a...)
}

// Infof 普通信息输出
func Infof(format string, args ...interface{}) {
	if logFile.level["Info"] {
		logFile.info.Println(fmt.Sprintf(format, args...))
	}
}

// Warnf 警告信息输出
func Warnf(format string, args ...interface{}) {
	if logFile.level["Warn"] {
		logFile.warn.Println(fmt.Sprintf(format, args...))
	}
}

// Debugf 调试信息输出
func Debugf(format string, args ...interface{}) {
	if logFile.level["Debug"] {
		logFile.debug.Println(fmt.Sprintf(format, args...))
	}
}

// Errorf 错误信息输出
func Errorf(format string, args ...interface{}) {
	if logFile.level["Error"] {
		logFile.err.Println(fmt.Sprintf(format, args...))
	}
}

// CreateLogFile 创建日志分割文件
func (m *file) CreateLogFile() {
	if logF != nil {
		logF.Close()
	}
	// 如目录不存在则创建
	CreateDir(logPath)
	var err error
	logFile.fileName = fmt.Sprintf("%s/%s%s", logPath, time.Now().Format("2006-01-02"), ".log")
	logF, err = os.OpenFile(logFile.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	logFile.fileFd = logF
	logFile.date = time.Now().Hour()
	logFile.info2 = log.New(io.MultiWriter(os.Stdout, logFile.fileFd), logTitle, log.Ldate|log.Ltime)
	logFile.info = log.New(io.MultiWriter(os.Stdout, logFile.fileFd), logTitle, log.Ldate|log.Ltime)
	logFile.warn = log.New(io.MultiWriter(os.Stdout, logFile.fileFd), logTitle, log.Ldate|log.Ltime)
	logFile.err = log.New(io.MultiWriter(os.Stderr, logFile.fileFd), logTitle, log.Ldate|log.Ltime)
	logFile.debug = log.New(io.MultiWriter(os.Stderr, logFile.fileFd), logTitle, log.Ldate|log.Ltime)
}
