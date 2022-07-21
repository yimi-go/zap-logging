package zap

import (
	"reflect"
	"testing"

	"go.uber.org/zap"

	"github.com/yimi-go/logging"
)

func TestNewFactory(t *testing.T) {
	type args struct {
		options *Options
	}
	tests := []struct {
		name     string
		args     args
		validate func(f *zapFactory, t *testing.T)
	}{
		{
			name: "nil",
			args: args{nil},
			validate: func(f *zapFactory, t *testing.T) {
				newOptions := NewOptions()
				if !reflect.DeepEqual(f.options.Load().(*Options), newOptions) {
					t.Errorf("NewZapFactory() options = %v, want %v", f.options.Load().(*Options), newOptions)
				}
				if f.zlCache.Load() == nil {
					t.Errorf("NewZapFactory() zlCache is nil")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFactory(tt.args.options)
			tt.validate(got, t)
		})
	}
}

func Test_zapFactory_Logger(t *testing.T) {
	factory := NewFactory(nil)
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		factory *zapFactory
		args    args
		want    logging.Logger
	}{
		{
			name:    "Logger",
			factory: factory,
			args: args{
				name: "test",
			},
			want: &zapLogger{
				fields:  nil,
				name:    "test",
				factory: factory,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.factory.Logger(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zapFactory_zap(t *testing.T) {
	factory := NewFactory(nil)
	c1 := factory.zlCache.Load()
	root := factory.zap("")
	if root == nil {
		t.Errorf("zapFactory.zap() root is nil")
	}
	c2 := factory.zlCache.Load()
	if reflect.DeepEqual(c1, c2) {
		t.Errorf("zapFactory.zap() zlCache is not updated")
	}
	root2 := factory.zap("")
	if root != root2 {
		t.Errorf("zapFactory.zap() root is not equal")
	}
	c3 := factory.zlCache.Load()
	if !reflect.DeepEqual(c2, c3) {
		t.Errorf("zapFactory.zap() zlCache is updated")
	}
	foo := factory.zap("foo")
	if foo == nil {
		t.Errorf("zapFactory.zap() foo is nil")
	}
}

func Test_zapFactory_SwitchOptions(t *testing.T) {
	factory := NewFactory(nil)
	o1 := factory.options.Load()
	_ = factory.zap("")
	c1 := factory.zlCache.Load()
	factory.SwitchOptions(nil)
	o2 := factory.options.Load()
	c2 := factory.zlCache.Load()
	if o1 != o2 {
		t.Errorf("zapFactory.SwitchOptions(nil) changed options")
	}
	if !reflect.DeepEqual(c1, c2) {
		t.Errorf("zapFactory.SwitchOptions(nil) changed zlCache")
	}
	factory.SwitchOptions(&Options{})
	o3 := factory.options.Load()
	c3 := factory.zlCache.Load()
	if o1 == o3 {
		t.Errorf("zapFactory.SwitchOptions() options is not updated")
	}
	if reflect.DeepEqual(c1, c3) {
		t.Errorf("zapFactory.SwitchOptions() zlCache is not updated")
	}
	if len(c3.(map[string]*zap.Logger)) != len(c2.(map[string]*zap.Logger)) {
		t.Errorf("zapFactory.SwitchOptions() zlCache size should not change")
	}
}
