package loggers

import (
	"snapp/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)


type zapLogger struct {
	cfg *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	l := &zapLogger{cfg: cfg}
	l.Init()
	return l
}

func (l *zapLogger) Init() {
	cfg := config.GetConfig()
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: cfg.Logger.Path,
		MaxSize: 10,
		MaxAge: 10,
		Compress: true,
		MaxBackups: 10,
		LocalTime: true,
	})

	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(cfg),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	l.logger = logger
}

var LogLevels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func (l zapLogger) getLogLevel(cfg *config.Config) zapcore.Level {
	level, exist := LogLevels[cfg.Logger.LogLevel]
	if !exist {
		return zap.DebugLevel
	}
	return level
}

func MapToParam(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	var param []interface{}
	for k, v := range extra {
		param = append(param, string(k))
		param = append(param, v)
	}
	return param
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := MapToParam(cat, sub, extra)
	l.logger.Debugw(msg, param...)
}
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := MapToParam(cat, sub, extra)
	l.logger.Infow(msg, param...)
}
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := MapToParam(cat, sub, extra)
	l.logger.Warnw(msg, param...)
}
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := MapToParam(cat, sub, extra)
	l.logger.Errorw(msg, param...)
}
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	param := MapToParam(cat, sub, extra)
	l.logger.Fatalw(msg, param...)
}
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}