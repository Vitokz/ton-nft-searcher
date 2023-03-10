package logger

import (
	"io"

	"github.com/rs/zerolog"
)

// Option defines logger option contract.
type Option interface {
	apply(config) config
}

// WithModuleName defines logger option to set module name.
func WithModuleName(name string) Option { //nolint:ireturn
	return loggerOptionFunc(func(cfg config) config {
		cfg.moduleName = name
		return cfg
	})
}

// WithAsyncMode defines logger option to run logger in fast async mode.
func WithAsyncMode() Option { //nolint:ireturn
	return loggerOptionFunc(func(cfg config) config {
		cfg.asyncMode = true
		return cfg
	})
}

// WithDiscardWriter defines logger option to run logger with discard writer.
func WithDiscardWriter() Option { //nolint:ireturn
	return WithConsoleWriter(io.Discard)
}

// WithConsoleWriter defines logger option to run logger with specified writer.
func WithConsoleWriter(writer io.Writer, levels ...zerolog.Level) Option { //nolint:ireturn
	return loggerOptionFunc(func(cfg config) config {
		if len(levels) == 0 {
			consoleWriter := zerolog.ConsoleWriter{Out: writer, TimeFormat: defaultTimeFormat}
			writer = configureOutputMessage(consoleWriter)
		}

		writer = &FilteredWriter{writer, levels}

		cfg.writers = append(cfg.writers, writer)

		return cfg
	})
}

// WithDefaultStderrLevels define default error logger levels.
func WithDefaultStderrLevels() []zerolog.Level {
	return []zerolog.Level{zerolog.ErrorLevel, zerolog.PanicLevel, zerolog.FatalLevel}
}

type config struct {
	moduleName string
	asyncMode  bool
	writers    []io.Writer
}

type loggerOptionFunc func(config) config

func (fn loggerOptionFunc) apply(cfg config) config {
	return fn(cfg)
}
