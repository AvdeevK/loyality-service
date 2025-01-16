package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Создание указателя на "пустой" логгер, который игнорирует все сообщения.
var Log *zap.Logger = zap.NewNop()

// Функция инициализации принимает уровень логирования - с уровня которого
// и выше будут выводиться все логи сервиса. В функции кастомные настройки логгера.
func Init(level string) error {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "message",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	cfg := zap.NewProductionConfig()
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		panic(err)
	}
	cfg.Level = lvl
	//cfg.OutputPaths = []string{"stdout"}
	cfg.Encoding = "console"
	cfg.EncoderConfig = encoderConfig
	Log, err = cfg.Build()

	return err
}
