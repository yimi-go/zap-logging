package zap

import (
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/yimi-go/logging"
)

type zapFactory struct {
	options atomic.Value
	zlCache atomic.Value
}

func NewFactory(options *Options) *zapFactory {
	if options == nil {
		options = NewOptions()
	}
	options = options.Defaulted()
	zf := &zapFactory{}
	zf.options.Store(options)
	zf.zlCache.Store(map[string]*zap.Logger{})
	return zf
}

func (z *zapFactory) Logger(name string) logging.Logger {
	return &zapLogger{
		name:    name,
		factory: z,
	}
}

func (z *zapFactory) level(name string) logging.Level {
	options := z.options.Load().(*Options)
	return options.level(name)
}

func (z *zapFactory) zap(name string) *zap.Logger {
	zc := z.zlCache.Load().(map[string]*zap.Logger)
	zl, ok := zc[name]
	if ok {
		return zl
	}
	zl = z.options.Load().(*Options).newZapLogger(name)
	newZc := make(map[string]*zap.Logger, len(zc)+1)
	for k, v := range zc {
		newZc[k] = v
	}
	newZc[name] = zl
	z.zlCache.Store(newZc)
	return zl
}

func (z *zapFactory) SwitchOptions(options *Options) {
	if options == nil {
		return
	}
	options = options.Defaulted()
	zc := z.zlCache.Load().(map[string]*zap.Logger)
	newZc := make(map[string]*zap.Logger, len(zc))
	for name := range zc {
		newZc[name] = options.newZapLogger(name)
	}
	z.options.Store(options)
	z.zlCache.Store(newZc)
}
