package configs

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

// Init initializes the global logger instance with a custom configuration.
func Init() {
	once.Do(func() {
		config := zap.Config{
			OutputPaths: []string{"stdout"},
			Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:   "msg",
				LevelKey:     "level",
				TimeKey:      "ts",
				NameKey:      "logger",
				CallerKey:    "caller",
				EncodeTime:   zapcore.ISO8601TimeEncoder,
				EncodeLevel:  zapcore.LowercaseLevelEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		var err error
		log, err = config.Build()
		if err != nil {
			panic(err)
		}
	})
}

// Info logs a message at InfoLevel.
func Info(message string, fields ...zap.Field) {
	ensureInitialized()
	log.Info(message, fields...)
}

// Error logs a message at ErrorLevel.
func Error(message string, err error, fields ...zap.Field) {
	ensureInitialized()
	fields = append(fields, zap.Error(err))
	log.Error(message, fields...)
}

// Fatal logs a message at FatalLevel and then calls os.Exit(1).
func Fatal(message string, err error, fields ...zap.Field) {
	ensureInitialized()
	fields = append(fields, zap.Error(err))
	log.Fatal(message, fields...)
}

// Debug logs a message at DebugLevel.
func Debug(message string, fields ...zap.Field) {
	ensureInitialized()
	log.Debug(message, fields...)
}

// Warn logs a message at WarnLevel.
func Warn(message string, fields ...zap.Field) {
	ensureInitialized()
	log.Warn(message, fields...)
}

// Sync flushes any buffered log entries.
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}

func ensureInitialized() {
	if log == nil {
		Init()
	}
}
