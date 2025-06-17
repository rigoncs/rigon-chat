package zlog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"rigon-chat-server/internal/config"
	"runtime"
)

var logger *zap.Logger
var logPath string

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	conf := config.GetConfig()
	logPath = conf.LogPath
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	fileWriteSyncer := zapcore.AddSync(file)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriteSyncer, zap.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
	)
	logger = zap.New(core)
}

func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100, // 单个文件最大100M
		MaxBackups: 60,  // 最多保存60个文件，多余60个日志文件后，清理较旧的日志
		MaxAge:     1,   // 一天一切割
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getCallerInfoForLog 获取调用日志的函数信息,包括函数名、文件名、行号
func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用房的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) // Base 函数返回路径的最后一个元素，只保留函数名
	callerFields = append(callerFields, zap.String("file", file), zap.Int("line", line), zap.String("func", funcName))
	return
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}
