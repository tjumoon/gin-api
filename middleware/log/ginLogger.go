package log

import (
	"github.com/gin-gonic/gin"
	"gin-api/utils/config"
	"log"
	"time"
	logUtil "gin-api/utils/log"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

func GetContextLogInfo(c *gin.Context) (string, int, string, string, string){
	method := c.Request.Method
	statusCode := c.Writer.Status()
	urlPath := c.Request.URL.Path
	errorString := c.Errors.String()
	clientIP := c.ClientIP()
	return method, statusCode, urlPath, errorString, clientIP
}

func AccessLogger() gin.HandlerFunc {
	out := logUtil.LumberJackLogger(config.AccessLogFilePath+ "-" + time.Now().Format("2006-01-02") +config.AccessLogFileExtension,
		config.AccessLogMaxSize, config.AccessLogMaxBackups, config.AccessLogMaxAge)
	stdLogger := log.New(out, "", 0)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		method, statusCode, urlPath, errorString, clientIP := GetContextLogInfo(c)

		stdLogger.Printf("[GIN] %v |%3d| %12v |%s %-7s | %s | %s",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			method,
			urlPath,
			clientIP,
			errorString,
		)
	}
}

func colorForStatus(status int) string {
	switch {
	case status >= 200 && status <= 299:
		return green
	case status >= 300 && status <= 399:
		return white
	case status >= 400 && status <= 499:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return blue
	case method == "POST":
		return cyan
	case method == "PUT":
		return yellow
	case method == "DELETE":
		return red
	case method == "PATCH":
		return green
	case method == "HEAD":
		return magenta
	case method == "OPTIONS":
		return white
	default:
		return reset
	}
}
