package setuplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// InitLogger инициализирует и возвращает zap-логгер
func InitLogger(env string) *zap.Logger {
	// Настройка кодировщика с подсветкой
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Цветной вывод уровня логов
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // Формат времени
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Создание кодировщика
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Настройка вывода в консоль
	consoleOutput := zapcore.Lock(os.Stdout)

	const (
		envLocal = "local"
		envDev   = "dev"
		envProd  = "prod"
	)

	var core zapcore.Core

	// Уровень логирования
	switch env {
	case envLocal:
		core = zapcore.NewCore(consoleEncoder, consoleOutput, zapcore.DebugLevel)
	case envDev:
		core = zapcore.NewCore(consoleEncoder, consoleOutput, zapcore.DebugLevel)
	case envProd:
		core = zapcore.NewCore(consoleEncoder, consoleOutput, zapcore.InfoLevel)
	default:
		core = zapcore.NewCore(consoleEncoder, consoleOutput, zapcore.InfoLevel)

	}

	// Создание логгера
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}
