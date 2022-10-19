package zap

import (
	"strings"

	"github.com/yimi-go/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FieldKeys struct {
	// Time is log time field name. "ts" as default.
	Time string `json:"time,omitempty"       yaml:"time,omitempty"`
	// Level is log level field name. "level" as default.
	Level string `json:"level,omitempty"      yaml:"level,omitempty"`
	// Logger is log logger name field name. "logger" as default.
	Logger string `json:"logger,omitempty"     yaml:"logger,omitempty"`
	// Caller is log caller field name. "caller" as default.
	Caller string `json:"caller,omitempty"     yaml:"caller,omitempty"`
	// Message is log message field name. "msg" as default.
	Message string `json:"message,omitempty"    yaml:"message,omitempty"`
	// Stacktrace is log stacktrace field name. "stacktrace" as default.
	Stacktrace string `json:"stacktrace,omitempty" yaml:"stacktrace,omitempty"`
}

// Defaulted returns a new FieldKeys filling blank items with default values.
func (f FieldKeys) Defaulted() FieldKeys {
	res := FieldKeys{
		Time:       "ts",
		Level:      "level",
		Logger:     "logger",
		Caller:     "caller",
		Message:    "msg",
		Stacktrace: "stacktrace",
	}
	time := strings.TrimSpace(f.Time)
	if len(time) != 0 {
		res.Time = time
	}
	level := strings.TrimSpace(f.Level)
	if len(level) != 0 {
		res.Level = level
	}
	logger := strings.TrimSpace(f.Logger)
	if len(logger) != 0 {
		res.Logger = logger
	}
	caller := strings.TrimSpace(f.Caller)
	if len(caller) != 0 {
		res.Caller = caller
	}
	message := strings.TrimSpace(f.Message)
	if len(message) != 0 {
		res.Message = message
	}
	stacktrace := strings.TrimSpace(f.Stacktrace)
	if len(stacktrace) != 0 {
		res.Stacktrace = stacktrace
	}
	return res
}

// Options is zap logger options.
type Options struct {
	// Levels is the minimum enabled logging levels, mapped by logger names.
	Levels map[string]logging.Level `json:"levels,omitempty" yaml:"levels,omitempty,flow"`
	// AddCallerSkipAdjusts is the adjustment for adjusting caller skips of caller annotation of specific logger.
	AddCallerSkipAdjusts map[string]int `json:"add_caller_skip_adjusts,omitempty" yaml:"add_caller_skip_adjusts,omitempty"`
	// FieldKeys is names of fixed log globalFields.
	FieldKeys FieldKeys `json:"field_keys,omitempty" yaml:"field_keys,omitempty"`
	// TimeLayout is log time field formatting layout. "2006-01-02 15:04:05.000" as default.
	TimeLayout string `json:"time_layout,omitempty" yaml:"time_layout,omitempty"`
	// OutputPaths is user log output paths. ["stdout"] as default.
	OutputPaths []string `json:"output_paths,omitempty" yaml:"output_paths,omitempty,flow"`
	// ErrorOutputPaths is log's error output path. ["stderr"] as default.
	ErrorOutputPaths []string `json:"error_output_paths,omitempty" yaml:"error_output_paths,omitempty,flow"`
	globalFields     []logging.Field
	// GlobalAddCallerSkipAdjust is the global adjustment for adjusting caller skips of caller annotation.
	// This effects all loggers.
	GlobalAddCallerSkipAdjust int `json:"global_add_caller_skip_adjust,omitempty" yaml:"global_add_caller_skip_adjust,omitempty"`
	// Development indicates if we are in development environment. False as default.
	Development bool `json:"development,omitempty" yaml:"development,omitempty"`
	// DisableCaller indicates whether disable log caller field. False as default.
	DisableCaller bool `json:"disable_caller,omitempty" yaml:"disable_caller,omitempty"`
	// DisableStacktrace indicates whether disable stacktrace field of error level logs. False as default.
	DisableStacktrace bool `json:"disable_stacktrace,omitempty" yaml:"disable_stacktrace,omitempty"`
	// DisableLogger indicates whether disable logger field. False as default.
	DisableLogger bool `json:"disable_logger,omitempty" yaml:"disable_logger,omitempty"`
}

// Defaulted returns a new Options filling blank items with default values.
func (o *Options) Defaulted() *Options {
	res := &Options{
		Levels: map[string]logging.Level{
			"": logging.InfoLevel,
		},
		Development:               o.Development,
		TimeLayout:                "2006-01-02 15:04:05.000",
		DisableCaller:             o.DisableCaller,
		DisableStacktrace:         o.DisableStacktrace,
		DisableLogger:             o.DisableLogger,
		OutputPaths:               []string{"stdout"},
		ErrorOutputPaths:          []string{"stderr"},
		GlobalAddCallerSkipAdjust: o.GlobalAddCallerSkipAdjust,
		AddCallerSkipAdjusts:      map[string]int{},
		globalFields:              o.globalFields,
	}
	for name, level := range o.Levels {
		res.Levels[name] = level
	}
	timeLayout := strings.TrimSpace(o.TimeLayout)
	if len(timeLayout) != 0 {
		res.TimeLayout = timeLayout
	}
	res.FieldKeys = o.FieldKeys.Defaulted()
	outputPaths := make([]string, 0, len(o.OutputPaths))
	for _, path := range o.OutputPaths {
		path = strings.TrimSpace(path)
		if len(path) != 0 {
			outputPaths = append(outputPaths, path)
		}
	}
	if len(outputPaths) != 0 {
		res.OutputPaths = outputPaths
	}
	errorOutputPaths := make([]string, 0, len(o.ErrorOutputPaths))
	for _, path := range o.ErrorOutputPaths {
		path = strings.TrimSpace(path)
		if len(path) != 0 {
			errorOutputPaths = append(errorOutputPaths, path)
		}
	}
	if len(errorOutputPaths) != 0 {
		res.ErrorOutputPaths = errorOutputPaths
	}
	for name, adj := range o.AddCallerSkipAdjusts {
		res.AddCallerSkipAdjusts[name] = adj
	}
	return res
}

// Option is zap-logging config option func.
type Option func(o *Options)

// Levels returns an Option that set log levels.
//
// If the parameter is empty, the default value would be used, which is {"":InfoLevel}
func Levels(levels map[string]logging.Level) Option {
	return func(o *Options) {
		o.Levels = levels
	}
}

// Development returns an Option that set whether we use development profile.
func Development(dev bool) Option {
	return func(o *Options) {
		o.Development = dev
	}
}

// TimeLayout returns an Option that set time field formatting layout.
//
// If the parameter is empty, the default value would be used, which is "2006-01-02 15:04:05.000".
func TimeLayout(layout string) Option {
	return func(o *Options) {
		o.TimeLayout = layout
	}
}

// TimeFieldKey returns an Option that set time field name.
//
// If the parameter is empty, the default value would be used, which is "ts".
func TimeFieldKey(key string) Option {
	return func(o *Options) {
		o.FieldKeys.Time = key
	}
}

// LevelFieldKey returns an Option that set level field name.
//
// If the parameter is empty, the default value would be used, which is "level".
func LevelFieldKey(key string) Option {
	return func(o *Options) {
		o.FieldKeys.Level = key
	}
}

// LoggerFieldKey returns an Option that set logger name field name.
//
// If the parameter is empty, the default value would be used, which is "logger".
func LoggerFieldKey(key string) Option {
	return func(o *Options) {
		o.FieldKeys.Logger = key
	}
}

// CallerFieldKey returns an Option that set caller field name.
//
// If the parameter is empty, the default value would be used, which is "caller".
func CallerFieldKey(key string) Option {
	return func(o *Options) {
		o.FieldKeys.Caller = key
	}
}

// MessageFieldKey returns an Option that set message field name.
//
// If the parameter is empty, the default value would be used, which is "msg".
func MessageFieldKey(key string) Option {
	return func(o *Options) {
		o.FieldKeys.Message = key
	}
}

// StacktraceFieldKey returns an Option that set stacktrace field name.
//
// If the parameter is empty, the default value would be used, which is "stacktrace".
func StacktraceFieldKey(key string) Option {
	return func(o *Options) {
		o.FieldKeys.Stacktrace = key
	}
}

// DisableCaller returns an Option that set whether disable caller field.
func DisableCaller(disable bool) Option {
	return func(o *Options) {
		o.DisableCaller = disable
	}
}

// DisableStacktrace returns an Option that set whether disable stacktrace field of error level logs.
func DisableStacktrace(disable bool) Option {
	return func(o *Options) {
		o.DisableStacktrace = disable
	}
}

// DisableLogger returns an Option that set whether disable logger field.
func DisableLogger(disable bool) Option {
	return func(o *Options) {
		o.DisableLogger = disable
	}
}

// OutputPaths returns an Option that set user log output paths.
//
// If the parameters are empty, the default value would be used, which is ["stdout"].
func OutputPaths(path ...string) Option {
	return func(o *Options) {
		o.OutputPaths = path
	}
}

// ErrorOutputPaths returns an Option that sets log's error output paths.
//
// If the parameters are empty, the default value would be used, which is ["stderr"].
func ErrorOutputPaths(path ...string) Option {
	return func(o *Options) {
		o.ErrorOutputPaths = path
	}
}

// GlobalAddCallerSkipAdjust returns an Option that sets global caller skips adjustment.
func GlobalAddCallerSkipAdjust(adjustment int) Option {
	return func(o *Options) {
		o.GlobalAddCallerSkipAdjust = adjustment
	}
}

// AddCallerSkipAdjust returns an Option that sets caller skips adjustment of logger name.
func AddCallerSkipAdjust(name string, adjustment int) Option {
	return func(o *Options) {
		o.AddCallerSkipAdjusts[name] = adjustment
	}
}

// GlobalFields returns an Option that sets global preset log fields.
func GlobalFields(fields ...logging.Field) Option {
	return func(o *Options) {
		o.globalFields = fields
	}
}

// NewOptions creates Options.
func NewOptions(options ...Option) *Options {
	res := &Options{}
	res = res.Defaulted()
	for _, o := range options {
		o(res)
	}
	return res
}

func (o *Options) newZapLogger(name string) *zap.Logger {
	levelEncoder := zapcore.CapitalColorLevelEncoder
	encoding := "console"
	if !o.Development {
		levelEncoder = zapcore.CapitalLevelEncoder
		encoding = "json"
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     o.FieldKeys.Message,
		LevelKey:       o.FieldKeys.Level,
		TimeKey:        o.FieldKeys.Time,
		NameKey:        o.FieldKeys.Logger,
		CallerKey:      o.FieldKeys.Caller,
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  o.FieldKeys.Stacktrace,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    levelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(o.TimeLayout),
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       o.Development,
		DisableCaller:     o.DisableCaller,
		DisableStacktrace: o.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      o.OutputPaths,
		ErrorOutputPaths: o.ErrorOutputPaths,
	}

	l, _ := loggerConfig.Build(
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCallerSkip(1),
		zap.AddCallerSkip(o.GlobalAddCallerSkipAdjust),
		zap.AddCallerSkip(o.AddCallerSkipAdjusts[name]),
	)
	if !o.DisableLogger {
		name = strings.TrimSpace(name)
		l = l.Named(name)
	}
	return l
}

func (o *Options) level(name string) logging.Level {
	name = strings.TrimSpace(name)
	if level, ok := o.Levels[name]; ok {
		return level
	}
	li := strings.LastIndexAny(name, "./:")
	if li == -1 {
		if rootLevel, ok := o.Levels[""]; ok {
			return rootLevel
		}
		if rootLevel, ok := o.Levels["root"]; ok {
			return rootLevel
		}
		return logging.InfoLevel
	}
	return o.level(name[:li])
}
