package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	l *zerolog.Logger
}

var L *Logger

func Init(level string) {
	var lvl zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		lvl = zerolog.ErrorLevel
	case "warn":
		lvl = zerolog.WarnLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "debug":
		lvl = zerolog.DebugLevel
	default:
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(lvl)

	// skipFrameCount := 3
	// l := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()

	L = &Logger{
		l: &l,
	}
}

func (lg *Logger) Debug(message interface{}, args ...interface{}) {
	lg.msg("debug", message, args...)
}

func (lg *Logger) Info(message string, args ...interface{}) {
	lg.log(message, args...)
}

func (lg *Logger) Warn(message string, args ...interface{}) {
	lg.log(message, args...)
}

func (lg *Logger) Error(message interface{}, args ...interface{}) {
	if lg.l.GetLevel() == zerolog.DebugLevel {
		lg.Debug(message, args...)
	}

	lg.msg("error", message, args...)
}

func (lg *Logger) ErrorStackTrace(err error) {
	lg.l.Error().Stack().Err(err).Msg("")
}

func (lg *Logger) Fatal(message interface{}, args ...interface{}) {
	lg.msg("fatal", message, args...)
	os.Exit(1)
}

func (lg *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		lg.l.Info().Msg(message)
	} else {
		lg.l.Info().Msgf(message, args...)
	}
}

func (lg *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		lg.log(msg.Error(), args...)
	case string:
		lg.log(msg, args...)
	default:
		lg.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
