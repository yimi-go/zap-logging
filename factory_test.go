package zap

import (
	"reflect"
	"testing"

	"github.com/yimi-go/logging"
)

func TestNewFactory(t *testing.T) {
	type args struct {
		options *Options
	}
	tests := []struct {
		args     args
		validate func(f *zapFactory, t *testing.T)
		name     string
	}{
		{
			name: "nil",
			args: args{nil},
			validate: func(f *zapFactory, t *testing.T) {
				newOptions := NewOptions()
				if !reflect.DeepEqual(f.options.Load().(*Options), newOptions) {
					t.Errorf("NewZapFactory() options = %v, want %v", f.options.Load().(*Options), newOptions)
				}
				if f.zlCache == nil {
					t.Errorf("NewZapFactory() zlCache is nil")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFactory(tt.args.options).(*zapFactory)
			tt.validate(got, t)
		})
	}
}

func Test_zapFactory_Logger(t *testing.T) {
	field := logging.String("foo", "bar")
	factory := NewFactory(NewOptions(GlobalFields(field))).(*zapFactory)
	type args struct {
		name string
	}
	tests := []struct {
		want    logging.Logger
		factory *zapFactory
		name    string
		args    args
	}{
		{
			name:    "Logger",
			factory: factory,
			args: args{
				name: "test",
			},
			want: &zapLogger{
				fields:  []logging.Field{field},
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
	factory := NewFactory(nil).(*zapFactory)
	root := factory.zap("")
	if root == nil {
		t.Errorf("zapFactory.zap() root is nil")
	}
	root2 := factory.zap("")
	if root != root2 {
		t.Errorf("zapFactory.zap() root is not equal")
	}
	foo := factory.zap("foo")
	if foo == nil {
		t.Errorf("zapFactory.zap() foo is nil")
	}
}

func Test_zapFactory_SwitchOptions(t *testing.T) {
	factory := NewFactory(nil).(*zapFactory)
	o1 := factory.options.Load()
	_ = factory.zap("")
	factory.SwitchOptions(nil)
	o2 := factory.options.Load()
	if o1 != o2 {
		t.Errorf("zapFactory.SwitchOptions(nil) changed options")
	}
	factory.SwitchOptions(&Options{})
	o3 := factory.options.Load()
	if o1 == o3 {
		t.Errorf("zapFactory.SwitchOptions() options is not updated")
	}
}
