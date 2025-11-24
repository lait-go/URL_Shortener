package logger

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

type Logger interface {
	InfoContext(ctx context.Context, msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, err error, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	With(args ...any) Logger
}

type Config struct {
	Level      string `envconfig:"LOGGER_LEVEL" default:"info"`
	JsonFormat bool   `envconfig:"LOGGER_JSON_FORMAT" default:"false"`
}

type slogLogger struct {
	logger *slog.Logger
	mu     sync.Mutex
}

var (
	defaultLogger Logger
	once          sync.Once
)

func Init(cfg Config) {
	once.Do(func() {
		var level slog.Level
		switch cfg.Level {
		case "debug":
			level = slog.LevelDebug
		case "info":
			level = slog.LevelInfo
		case "warn":
			level = slog.LevelWarn
		case "error":
			level = slog.LevelError
		default:
			level = slog.LevelInfo
		}

		var handler slog.Handler
		opts := &slog.HandlerOptions{
			Level: level,
		}

		if cfg.JsonFormat {
			handler = slog.NewJSONHandler(os.Stdout, opts)
		} else {
			handler = slog.NewTextHandler(os.Stdout, opts)
		}

		l := slog.New(handler)
		defaultLogger = &slogLogger{logger: l}
		
		// Set as default logger for stdlib
		slog.SetDefault(l)
	})
}

func Default() Logger {
	if defaultLogger == nil {
		// Fallback to basic logger if not initialized
		Init(Config{Level: "info", JsonFormat: false})
	}
	return defaultLogger
}

func (l *slogLogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *slogLogger) ErrorContext(ctx context.Context, msg string, err error, args ...any) {
	allArgs := make([]any, 0, len(args)+2)
	allArgs = append(allArgs, args...)
	allArgs = append(allArgs, slog.String("error", err.Error()))
	l.logger.ErrorContext(ctx, msg, allArgs...)
}

func (l *slogLogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *slogLogger) With(args ...any) Logger {
	return &slogLogger{logger: l.logger.With(args...)}
}