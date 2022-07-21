package zap

import (
	"fmt"

	"go.uber.org/zap/zapcore"

	"github.com/yimi-go/logging"
)

type zapLogger struct {
	name    string
	factory *zapFactory
	fields  []logging.Field
}

func sprintln(v ...any) string {
	s := fmt.Sprintln(v...)
	return s[:len(s)-1]
}

func (z *zapLogger) Enabled(level logging.Level) bool {
	return z.factory.level(z.name).Enabled(level)
}

func (z *zapLogger) Debug(v ...any) {
	if !z.Enabled(logging.DebugLevel) {
		return
	}
	z.factory.zap(z.name).Debug(fmt.Sprint(v...), z.zapFields()...)
}

func (z *zapLogger) Debugln(v ...any) {
	if !z.Enabled(logging.DebugLevel) {
		return
	}
	z.factory.zap(z.name).Debug(sprintln(v...), z.zapFields()...)
}

func (z *zapLogger) Debugf(format string, v ...any) {
	if !z.Enabled(logging.DebugLevel) {
		return
	}
	z.factory.zap(z.name).Debug(fmt.Sprintf(format, v...), z.zapFields()...)
}

func (z *zapLogger) Debugw(message string, field ...logging.Field) {
	if !z.Enabled(logging.DebugLevel) {
		return
	}
	z.factory.zap(z.name).Debug(message, z.zapFields(field...)...)
}

func (z *zapLogger) Info(v ...any) {
	if !z.Enabled(logging.InfoLevel) {
		return
	}
	z.factory.zap(z.name).Info(fmt.Sprint(v...), z.zapFields()...)
}

func (z *zapLogger) Infoln(v ...any) {
	if !z.Enabled(logging.InfoLevel) {
		return
	}
	z.factory.zap(z.name).Info(sprintln(v...), z.zapFields()...)
}

func (z *zapLogger) Infof(format string, v ...any) {
	if !z.Enabled(logging.InfoLevel) {
		return
	}
	z.factory.zap(z.name).Info(fmt.Sprintf(format, v...), z.zapFields()...)
}

func (z *zapLogger) Infow(message string, field ...logging.Field) {
	if !z.Enabled(logging.InfoLevel) {
		return
	}
	z.factory.zap(z.name).Info(message, z.zapFields(field...)...)
}

func (z *zapLogger) Warn(v ...any) {
	if !z.Enabled(logging.WarnLevel) {
		return
	}
	z.factory.zap(z.name).Warn(fmt.Sprint(v...), z.zapFields()...)
}

func (z *zapLogger) Warnln(v ...any) {
	if !z.Enabled(logging.WarnLevel) {
		return
	}
	z.factory.zap(z.name).Warn(sprintln(v...), z.zapFields()...)
}

func (z *zapLogger) Warnf(format string, v ...any) {
	if !z.Enabled(logging.WarnLevel) {
		return
	}
	z.factory.zap(z.name).Warn(fmt.Sprintf(format, v...), z.zapFields()...)
}

func (z *zapLogger) Warnw(message string, field ...logging.Field) {
	if !z.Enabled(logging.WarnLevel) {
		return
	}
	z.factory.zap(z.name).Warn(message, z.zapFields(field...)...)
}

func (z *zapLogger) Error(v ...any) {
	if !z.Enabled(logging.ErrorLevel) {
		return
	}
	z.factory.zap(z.name).Error(fmt.Sprint(v...), z.zapFields()...)
}

func (z *zapLogger) Errorln(v ...any) {
	if !z.Enabled(logging.ErrorLevel) {
		return
	}
	z.factory.zap(z.name).Error(sprintln(v...), z.zapFields()...)
}

func (z *zapLogger) Errorf(format string, v ...any) {
	if !z.Enabled(logging.ErrorLevel) {
		return
	}
	z.factory.zap(z.name).Error(fmt.Sprintf(format, v...), z.zapFields()...)
}

func (z *zapLogger) Errorw(message string, field ...logging.Field) {
	if !z.Enabled(logging.ErrorLevel) {
		return
	}
	z.factory.zap(z.name).Error(message, z.zapFields(field...)...)
}

func (z *zapLogger) WithField(field ...logging.Field) logging.Logger {
	return &zapLogger{
		name:    z.name,
		factory: z.factory,
		fields:  append(z.fields, field...),
	}
}

func (z *zapLogger) zapFields(field ...logging.Field) []zapcore.Field {
	fields := make([]zapcore.Field, 0, len(field)+len(z.fields))
	for _, f := range z.fields {
		fields = append(fields, mapZapField(f))
	}
	for _, f := range field {
		fields = append(fields, mapZapField(f))
	}
	return fields
}
