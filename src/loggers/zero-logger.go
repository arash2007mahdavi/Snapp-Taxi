package loggers

import (
	"os"
	"snapp/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type zeroLogger struct {
	cfg *config.Config
	logger *zerolog.Logger
}

func NewZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

var log_Levels = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info": zerolog.InfoLevel,
	"warn": zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exist := log_Levels[l.cfg.Logger.LogLevel]
	if exist {
		return level
	}
	return zerolog.DebugLevel
}

func (l *zeroLogger) Init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(l.getLogLevel())

	file, err := os.OpenFile(l.cfg.Logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(file).
		With().
		Timestamp().
		Str("AppName", "Snapp").
		Str("Logger", "ZeroLogger").
		Logger()

	l.logger = &logger
}

func extra_to_params(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extra_to_params(extra)

	l.logger.Debug().Fields(params).
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Msg(msg)
}
func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extra_to_params(extra)

	l.logger.Info().Fields(params).
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Msg(msg)
}
func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extra_to_params(extra)

	l.logger.Warn().Fields(params).
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Msg(msg)
}
func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extra_to_params(extra)

	l.logger.Error().Fields(params).
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Msg(msg)
}
func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := extra_to_params(extra)

	l.logger.Fatal().Fields(params).
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Msg(msg)
}
func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}