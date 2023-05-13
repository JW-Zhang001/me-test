package initialize

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewTee(
		// Write logs to both the console and the file
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
	)

	// zap.AddCaller()  Add the ability to log information about function calls
	logger := zap.New(core, zap.AddCaller())
	// Replace global logger
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Modified time coder

	// Log levels are recorded in uppercase letters in log files
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder Printing is more in line with the way people see
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	fileTime := fmt.Sprintf("logs/%s%s", time.Now().Format("2006-01-02 15:04:05"), ".log")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileTime,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

var Logger = zap.S()
