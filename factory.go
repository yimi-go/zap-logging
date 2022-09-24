package zap

import (
	"github.com/yimi-go/keeper"
	"github.com/yimi-go/logging"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type zapFactory struct {
	options atomic.Value
	zlCache keeper.Keeper[string, *zap.Logger]
}

func NewFactory(options *Options) logging.Factory {
	if options == nil {
		options = NewOptions()
	}
	options = options.Defaulted()
	zf := &zapFactory{}
	zf.options.Store(options)
	zf.zlCache = keeper.NewKeeper(func(key string) *zap.Logger {
		return zf.options.Load().(*Options).newZapLogger(key)
	})
	return zf
}

func (z *zapFactory) Logger(name string) logging.Logger {
	return &zapLogger{
		name:    name,
		factory: z,
		fields:  z.options.Load().(*Options).globalFields,
	}
}

func (z *zapFactory) level(name string) logging.Level {
	options := z.options.Load().(*Options)
	return options.level(name)
}

func (z *zapFactory) zap(name string) *zap.Logger {
	return z.zlCache.Get(name)
}

func (z *zapFactory) SwitchOptions(options *Options) {
	if options == nil {
		return
	}
	options = options.Defaulted()
	z.options.Store(options)
	z.zlCache.Clear()
}
