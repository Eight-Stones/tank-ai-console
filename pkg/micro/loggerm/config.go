package loggerm

import (
	"os"

	"go-micro.dev/v4/logger"
	"go.uber.org/zap/zapcore"
)

type Level logger.Level

const (
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel Level = iota - 2
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// InfoLevel is the default logging priority.
	// General operational entries about what's going on inside the application.
	InfoLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	ErrorLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. highest level of severity.
	FatalLevel
)

type Option func(*options)

type options struct {
	Level             Level
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string                 `json:"encoding" yaml:"encoding"`
	EncoderConfig     zapcore.EncoderConfig  `json:"encoderConfig" yaml:"encoderConfig"`
	OutputPaths       []string               `json:"outputPaths" yaml:"outputPaths"`
	ErrorOutputPaths  []string               `json:"errorOutputPaths" yaml:"errorOutputPaths"`
	InitialFields     map[string]interface{} `json:"initialFields" yaml:"initialFields"`
}

func WithLevelString(level string) Option {
	return func(opts *options) {
		mapping := map[string]Level{
			"trace": TraceLevel,
			"debug": DebugLevel,
			"info":  InfoLevel,
			"warn":  WarnLevel,
			"error": ErrorLevel,
			"fatal": FatalLevel,
		}
		opts.Level = mapping[level]
	}
}

// WithLevel set default level for the logger.
func WithLevel(level Level) Option {
	return func(opts *options) {
		opts.Level = level
	}
}

func WithDevelopment(in bool) Option {
	return func(opts *options) {
		opts.Development = in
	}
}

func WithDisableCaller(in bool) Option {
	return func(opts *options) {
		opts.DisableCaller = in
	}
}

func WithDisableStacktrace(in bool) Option {
	return func(opts *options) {
		opts.DisableStacktrace = in
	}
}

func WithEncoding(encoding string) Option {
	return func(opts *options) {
		opts.Encoding = encoding
	}
}

func WithOutputPaths(paths ...string) Option {
	return func(opts *options) {
		opts.OutputPaths = paths
	}
}

func WithErrorOutputPaths(paths ...string) Option {
	return func(opts *options) {
		opts.ErrorOutputPaths = paths
	}
}

func WithInitialFields(fields map[string]interface{}) Option {
	return func(opts *options) {
		for k, v := range fields {
			opts.InitialFields[k] = v
		}
	}
}

var defaultOptions = options{ //nolint:gochecknoglobals // only package variable
	Level:             DebugLevel,
	Development:       false,
	DisableCaller:     false,
	DisableStacktrace: false,
	Encoding:          "json",
	EncoderConfig: zapcore.EncoderConfig{
		TimeKey:    "timestamp",
		EncodeTime: zapcore.ISO8601TimeEncoder,
	},
	OutputPaths: []string{
		"stderr",
	},
	ErrorOutputPaths: []string{
		"stderr",
	},
	InitialFields: map[string]interface{}{
		"pid": os.Getpid(),
	},
}
