package loggerm

import (
	"os"

	"go-micro.dev/v4/logger"

	"go-micro-service-template/common"
)

const ExtraFieldsCount = 2

type Logger struct {
	Log logger.Logger
}

func Sugar(l logger.Logger) common.LoggerI {
	return &Logger{Log: l}
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...any) {
	l.Log.Log(logger.DebugLevel, args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...any) {
	l.Log.Logf(logger.DebugLevel, template, args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...any) {
	l.Log.Log(logger.InfoLevel, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...any) {
	l.Log.Logf(logger.InfoLevel, template, args...)
}

// Infow ...
func (l *Logger) Infow(kav ...any) {
	fields := make(map[string]any, len(kav)/2+1)
	for i := 0; i+1 < len(kav); i += 2 {
		fields[kav[i].(string)] = kav[i+1]
	}
	l.Log.Fields(fields).Log(logger.InfoLevel)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...any) {
	l.Log.Log(logger.WarnLevel, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...any) {
	l.Log.Logf(logger.WarnLevel, template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...any) {
	l.Log.Log(logger.ErrorLevel, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...any) {
	l.Log.Logf(logger.ErrorLevel, template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) Fatal(args ...any) {
	l.Log.Log(logger.FatalLevel, args...)
	os.Exit(1)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...any) {
	l.Log.Logf(logger.FatalLevel, template, args...)
	os.Exit(1)
}

// Print logs a message at level Debug on the ZapLogger.
func (l *Logger) Print(args ...any) {
	l.Log.Log(logger.DebugLevel, args...)
}

// Printf logs a message at level Debug on the ZapLogger.
func (l *Logger) Printf(template string, args ...any) {
	l.Log.Logf(logger.DebugLevel, template, args...)
}

func (l *Logger) Name(name string) common.LoggerI {
	return &Logger{Log: l.Log.Fields(map[string]any{"name": name})}
}

func (l *Logger) Fields(in ...any) common.LoggerI {
	fields := make(map[string]any, len(in)/ExtraFieldsCount)
	for i := 0; i < len(in) && i+1 < len(in); i += 2 {
		fields[in[i].(string)] = in[i+1]
	}
	return &Logger{Log: l.Log.Fields(fields)}
}
