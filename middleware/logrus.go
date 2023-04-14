package middleware

import (
	"bytes"
	"encoding/json"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-server/entity"
	"github.com/gin-server/global"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggerToFile() gin.HandlerFunc {
	logFilePath := global.GVA_VP.GetString("logs.path")
	logFileName := global.GVA_VP.GetString("logs.name")
	logFileMaxSize := global.GVA_VP.GetInt("logs.max_size")

	// // 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// // 【单文件写入】
	// src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }

	// 实例化
	logger := logrus.New()

	//設置輸出【单文件写入】
	// logger.SetOutput(os.Stdout)

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		//设置最大保存时间7天
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔1天
		rotatelogs.WithRotationTime(24*time.Hour),
		//文件达到多大则切割文件，单位为 bytes WithRotationTime and WithRotationSize 两者任意一个条件达到都会切割
		rotatelogs.WithRotationSize(int64(logFileMaxSize*1024)),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增钩子
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		var responseCode string
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			res := entity.Result{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = strconv.Itoa(res.Code)
				responseMsg = res.Message
				responseData = res.Data
			}
		}

		// 结束时间
		endTime := time.Now()

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}

		// 日志格式
		logger.WithFields(logrus.Fields{
			"request_method":    c.Request.Method,
			"request_uri":       c.Request.RequestURI,
			"request_proto":     c.Request.Proto,
			"request_useragent": c.Request.UserAgent(),
			"request_referer":   c.Request.Referer(),
			"request_post_data": c.Request.PostForm.Encode(),
			"request_client_ip": c.ClientIP(),

			"response_status_code": c.Writer.Status(),
			"response_code":        responseCode,
			"response_msg":         responseMsg,
			"response_data":        responseData,

			"cost_time": endTime.Sub(startTime),
		}).Info()
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
