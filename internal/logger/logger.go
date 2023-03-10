package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	red        = "\033[31m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	cyan       = "\033[36m"
	darkGray   = "\033[91m"
	colorReset = "\033[0m"

	moduleField = "module"

	defaultTimeFormat = time.RFC3339Nano
)

// New creates logger with specified options.
func New(opts ...Option) (zerolog.Logger, error) {
	var cfg config

	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}

	if cfg.writers == nil {
		return zerolog.Nop(), nil
	}

	if cfg.asyncMode {
		for i := range cfg.writers {
			if f, ok := cfg.writers[i].(*FilteredWriter); ok {
				f.AsyncWrap()
			}
		}
	}

	multiLevelWriter := zerolog.MultiLevelWriter(cfg.writers...)

	logger := zerolog.New(multiLevelWriter).With().Caller().Timestamp().Logger()

	if cfg.moduleName != "" {
		logger = logger.With().Str(moduleField, cfg.moduleName).Logger()
	}

	return logger, nil
}

// NewWithDefaultWriters creates logger with stdout and stderr writers.
func NewWithDefaultWriters(opts ...Option) (zerolog.Logger, error) {
	opts = append(opts, WithConsoleWriter(os.Stdout))
	opts = append(opts, WithConsoleWriter(os.Stderr, WithDefaultStderrLevels()...))

	return New(opts...)
}

func configureOutputMessage(writer zerolog.ConsoleWriter) zerolog.ConsoleWriter {
	// configure time format
	zerolog.TimeFieldFormat = time.RFC3339Nano
	writer.TimeFormat = "15:04:05.000000"

	// level color configure
	writer.FormatLevel = configureFormatLevel

	// cut message if its to long or add space in its to short
	writer.FormatMessage = configureFormatMessage

	writer.FormatFieldValue = func(i interface{}) string {
		if str, ok := i.(string); ok {
			return strings.ReplaceAll(str, `\"`, `"`)
		}

		return fmt.Sprintf("%s", i)
	}

	writer.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf(red+"%s=", i)
	}

	// formatting caller string
	zerolog.CallerMarshalFunc = configureCaller

	return writer
}

func configureFormatLevel(i interface{}) string {
	inputMessage, _ := i.(string)

	level, _ := zerolog.ParseLevel(inputMessage)

	switch level {
	case zerolog.DebugLevel:
		return fmt.Sprintf(cyan+"%-5s ➙"+colorReset, i)
	case zerolog.InfoLevel:
		return fmt.Sprintf(green+"%-5s ➙"+colorReset, i)
	case zerolog.WarnLevel:
		return fmt.Sprintf(yellow+"%-5s ➙"+colorReset, i)
	case zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel:
		return fmt.Sprintf(red+"%-5s ➙"+colorReset, i)
	case zerolog.NoLevel:
		return fmt.Sprintf(darkGray+"%-5s ➙"+colorReset, i)
	default:
		return fmt.Sprintf("%-5s ➙", i)
	}
}

const messagePostfix = "..."

func configureFormatMessage(i interface{}) string {
	// lengthRequired means how many symbols would be in message between code pointer and params
	// this was done for greater readability
	const lengthRequired = 80
	messageLength := len(fmt.Sprintf("%s", i))

	if messageLength <= lengthRequired {
		return fmt.Sprintf("%-*s", lengthRequired, i)
	}

	message := fmt.Sprintf("%s", i)
	// lengthRequired-3 it is 80 symbols string with ellipsis (80-3=77  len"..."=3)
	message = message[:lengthRequired-len(messagePostfix)] + messagePostfix

	return message
}

const (
	maxFileNameLen    = 8
	maxLinePointerLen = 4
	delimiter         = ":"
	callerPostfix     = "*"
)

func configureCaller(_ uintptr, file string, line int) string {
	// lengthRequired means how many symbols can contain caller info
	fileName := filepath.Base(file)
	linePointer := strconv.Itoa(line)

	if len(linePointer) > maxLinePointerLen {
		linePointer = linePointer[:maxLinePointerLen-len(callerPostfix)] + callerPostfix
	}

	restAvailableLen := (maxLinePointerLen - len(linePointer)) + maxFileNameLen
	if len(fileName) > restAvailableLen {
		fileName = fileName[:restAvailableLen-len(callerPostfix)] + callerPostfix
	}

	return fmt.Sprintf("%*s", maxFileNameLen+maxLinePointerLen+len(delimiter), fileName+delimiter+linePointer)
}

const (
	defaultWriterBufferSize   = 16 * 1024 * 1024 // 16 MB
	defaultWriterPollInterval = 10 * time.Millisecond
)
