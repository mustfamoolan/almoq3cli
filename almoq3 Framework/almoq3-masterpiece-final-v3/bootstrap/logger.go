package bootstrap

import (
	"os"
	"path/filepath"

	"almoq3-masterpiece-final-v3/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func InitializeLogger() {
	logConfig := config.Global.App
	
	// Ensure log directory exists
	logDir := "storage/logs"
	os.MkdirAll(logDir, 0755)
	
	logFile := filepath.Join(logDir, "app.log")

	// Lumberjack for log rotation
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	var level zapcore.Level
	if logConfig.Debug {
		level = zap.DebugLevel
	} else {
		level = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		level,
	)

	Log = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Log)

	Log.Info("Logger initialized successfully", zap.String("level", level.String()))
}
