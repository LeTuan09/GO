package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 3))
}

// format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encoderConfig.TimeKey = "time"
	// from info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// caller
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

/** 
 * * name: tên file
 * * flag: chế độ mở file: readOnly, writeOnly, ....
 * * perm: quyền đọc và ghi cho tất cả
*/ 
func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm) // name, flag, perm
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr) // luồng tiêu chuẩn
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

