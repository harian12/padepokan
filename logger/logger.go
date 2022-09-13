package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"padepokan/helpers"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logs *zap.Logger

// var sugarLogger *zap.SugaredLogger

var (
	dateNow = time.Now()
	config  = zap.NewProductionConfig()
	appName = helpers.GetEnv("APP_NAME", "Padepokan")
)

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.CallerKey = "function"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = TimeFormat
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.MessageKey = "message"
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.FunctionKey = "function key"
	config.EncoderConfig = encoderConfig
	err := os.MkdirAll("logger/data", os.ModePerm)
	err = os.MkdirAll("logger/data/error", os.ModePerm)
	err = os.MkdirAll("logger/data/debug", os.ModePerm)
	err = os.MkdirAll("logger/data/info", os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func Info(ctx *gin.Context, message string) {
	config.OutputPaths = []string{
		fmt.Sprintf("logger/data/info/%s-Info-%s.log", appName, dateNow.Format("2006-01-02")),
	}
	logs, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
	getHeader, _ := json.Marshal(ctx.Request.Header)
	header := json.RawMessage(string(getHeader)[:])
	getBody, _ := json.Marshal(ctx.Request.Body)
	body := json.RawMessage(string(getBody)[:])
	logs.Info(message, zap.Any("Header", &header), zap.Any("Body", body))
}

func Debug(ctx *gin.Context, message string, data interface{}) {
	config.OutputPaths = []string{
		fmt.Sprintf("logger/data/debug/%s-Debug-%s.log", appName, dateNow.Format("2006-01-02")),
	}
	logs, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
	getHeader, _ := json.Marshal(ctx.Request.Header)
	header := json.RawMessage(string(getHeader)[:])
	getBody, _ := json.Marshal(ctx.Request.Body)
	body := json.RawMessage(string(getBody)[:])

	resultjs, _ := json.Marshal(data)
	jsRaw := json.RawMessage(string(resultjs)[:])
	logs.Debug(message, zap.Any("debug", &jsRaw), zap.Any("Header", &header), zap.Any("Body", body))
}

func Error(ctx *gin.Context, message string, data interface{}) {
	config.OutputPaths = []string{
		fmt.Sprintf("logger/data/error/%s-Error-%s.log", appName, dateNow.Format("2006-01-02")),
	}
	logs, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	getHeader, _ := json.Marshal(ctx.Request.Header)
	header := json.RawMessage(string(getHeader)[:])
	getBody, _ := json.Marshal(ctx.Request.Body)
	body := json.RawMessage(string(getBody)[:])

	resultjs, _ := json.Marshal(data)
	jsRaw := json.RawMessage(string(resultjs)[:])
	logs.Error(message, zap.Any("error", &jsRaw), zap.Any("Header", &header), zap.Any("Body", body))
}

func ErrorFunc(message string, data interface{}) {
	config.OutputPaths = []string{
		fmt.Sprintf("logger/data/error/%s-Error-%s.log", appName, dateNow.Format("2006-01-02")),
	}
	logs, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
	resultjs, _ := json.Marshal(data)
	jsRaw := json.RawMessage(string(resultjs)[:])
	logs.Error(message, zap.Any("error", &jsRaw))
}

func iso3339CleanTime(time.Time) string {
	date := dateNow.Format("2006-01-02 15:04:05")
	return date
}

func TimeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(iso3339CleanTime(t))
}
