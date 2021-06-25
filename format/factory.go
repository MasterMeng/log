package format

import (
	"errors"
	"os"
	"path"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(logDir, format string) (*zap.Logger, error) {

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, err
		}
	}

	var coreArr []zapcore.Core

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	// log level
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	// info file
	infoFielWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(logDir, "info.log"),
		MaxSize:    2,
		MaxBackups: 100,
		MaxAge:     30,
		Compress:   false,
	})
	//error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(logDir, "error.log"), //日志文件存放目录
		MaxSize:    1,                              //文件大小限制,单位MB
		MaxBackups: 5,                              //最大保留日志文件数量
		MaxAge:     30,                             //日志文件保留天数
		Compress:   false,                          //是否压缩处理
	})

	var encoder zapcore.Encoder
	switch strings.ToLower(format) {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		return nil, errors.New("unknown format")
	}

	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFielWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)    //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	return zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()), nil
}
