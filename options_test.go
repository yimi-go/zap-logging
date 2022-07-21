package zap

import (
	"bufio"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yimi-go/logging"
)

func TestLevels(t *testing.T) {
	o := &Options{}
	levels := map[string]logging.Level{"foo": logging.InfoLevel}
	Levels(levels)(o)
	assert.Equal(t, levels, o.Levels)
}

func TestDevelopment(t *testing.T) {
	o := &Options{}
	Development(true)(o)
	assert.True(t, o.Development)
}

func TestTimeLayout(t *testing.T) {
	o := &Options{}
	layout := "aaa"
	TimeLayout(layout)(o)
	assert.Equal(t, layout, o.TimeLayout)
}

func TestTimeFieldKey(t *testing.T) {
	o := &Options{}
	key := "aaa"
	TimeFieldKey(key)(o)
	assert.Equal(t, key, o.FieldKeys.Time)
}

func TestLevelFieldKey(t *testing.T) {
	o := &Options{}
	key := "aaa"
	LevelFieldKey(key)(o)
	assert.Equal(t, key, o.FieldKeys.Level)
}

func TestLoggerFieldKey(t *testing.T) {
	o := &Options{}
	key := "aaa"
	LoggerFieldKey(key)(o)
	assert.Equal(t, key, o.FieldKeys.Logger)
}

func TestCallerFieldKey(t *testing.T) {
	o := &Options{}
	key := "aaa"
	CallerFieldKey(key)(o)
	assert.Equal(t, key, o.FieldKeys.Caller)
}

func TestMessageFieldKey(t *testing.T) {
	o := &Options{}
	key := "aaa"
	MessageFieldKey(key)(o)
	assert.Equal(t, key, o.FieldKeys.Message)
}

func TestStacktraceFieldKey(t *testing.T) {
	o := &Options{}
	key := "aaa"
	StacktraceFieldKey(key)(o)
	assert.Equal(t, key, o.FieldKeys.Stacktrace)
}

func TestDisableCaller(t *testing.T) {
	o := &Options{}
	DisableCaller(true)(o)
	assert.True(t, o.DisableCaller)
}

func TestDisableStacktrace(t *testing.T) {
	o := &Options{}
	DisableStacktrace(true)(o)
	assert.True(t, o.DisableStacktrace)
}

func TestDisableLogger(t *testing.T) {
	o := &Options{}
	DisableLogger(true)(o)
	assert.True(t, o.DisableLogger)
}

func TestOutputPaths(t *testing.T) {
	o := &Options{}
	path := []string{"a", "b"}
	OutputPaths(path...)(o)
	assert.Equal(t, path, o.OutputPaths)
}

func TestErrorOutputPaths(t *testing.T) {
	o := &Options{}
	path := []string{"a", "b"}
	ErrorOutputPaths(path...)(o)
	assert.Equal(t, path, o.ErrorOutputPaths)
}

func TestAddCallerSkipExtra(t *testing.T) {
	o := &Options{}
	extra := 2
	AddCallerSkipExtra(extra)(o)
	assert.Equal(t, extra, o.addCallerSkipExtra)
}

func TestGlobalFields(t *testing.T) {
	o := &Options{}
	fields := []logging.Field{
		logging.String("foo", "bar"),
	}
	GlobalFields(fields...)(o)
	assert.Equal(t, fields, o.globalFields)
}

func TestNewOptions(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Options
	}{
		{
			name: "nil_options",
			args: args{
				options: nil,
			},
			want: &Options{},
		},
		{
			name: "empty_options",
			args: args{
				options: []Option{},
			},
			want: &Options{
				Levels: map[string]logging.Level{
					"": logging.InfoLevel,
				},
			},
		},
		{
			name: "dev_options",
			args: args{
				options: []Option{
					func(options *Options) {
						options.Development = true
					},
				},
			},
			want: &Options{
				Levels: map[string]logging.Level{
					"": logging.InfoLevel,
				},
				Development: true,
			},
		},
		{
			name: "levels_options",
			args: args{
				options: []Option{
					func(options *Options) {
						options.Levels["foo"] = logging.DebugLevel
					},
				},
			},
			want: &Options{
				Levels: map[string]logging.Level{
					"":    logging.InfoLevel,
					"foo": logging.DebugLevel,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want.Defaulted()
			assert.Equal(t, want, NewOptions(tt.args.options...))
		})
	}
}

func TestOptions_level(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		options *Options
		args    args
		want    logging.Level
	}{
		{
			name: "nil_levels_root",
			options: &Options{
				Levels: nil,
			},
			args: args{name: ""},
			want: logging.InfoLevel,
		},
		{
			name: "nil_levels_foo",
			options: &Options{
				Levels: nil,
			},
			args: args{name: "foo"},
			want: logging.InfoLevel,
		},
		{
			name:    "default_levels_root",
			options: NewOptions(),
			args:    args{name: ""},
			want:    logging.InfoLevel,
		},
		{
			name:    "default_levels_foo",
			options: NewOptions(),
			args:    args{name: "foo"},
			want:    logging.InfoLevel,
		},
		{
			name: "root_debug_foo_info_root",
			options: &Options{
				Levels: map[string]logging.Level{
					"":    logging.DebugLevel,
					"foo": logging.InfoLevel,
				},
			},
			args: args{name: ""},
			want: logging.DebugLevel,
		},
		{
			name: "root_debug_foo_info_foo",
			options: &Options{
				Levels: map[string]logging.Level{
					"":    logging.DebugLevel,
					"foo": logging.InfoLevel,
				},
			},
			args: args{name: "foo"},
			want: logging.InfoLevel,
		},
		{
			name: "root_nil_foo_debug_root",
			options: &Options{
				Levels: map[string]logging.Level{
					"foo": logging.DebugLevel,
				},
			},
			args: args{name: ""},
			want: logging.InfoLevel,
		},
		{
			name: "root_nil_foo_debug_foo",
			options: &Options{
				Levels: map[string]logging.Level{
					"foo": logging.DebugLevel,
				},
			},
			args: args{name: "foo"},
			want: logging.DebugLevel,
		},
		{
			name: "root_level_2",
			options: &Options{
				Levels: map[string]logging.Level{
					"root": logging.ErrorLevel,
				},
			},
			args: args{name: ""},
			want: logging.ErrorLevel,
		},
		{
			name: "foo_bar",
			options: &Options{
				Levels: map[string]logging.Level{
					"foo": logging.DebugLevel,
				},
			},
			args: args{name: "foo.bar"},
			want: logging.DebugLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.options.level(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options.level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_newZapLogger(t *testing.T) {
	t.Run("prod", func(t *testing.T) {
		stdout := os.Stdout
		defer func() {
			os.Stdout = stdout
		}()
		r, w, err := os.Pipe()
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = r.Close()
		}()
		os.Stdout = w
		options := NewOptions(func(options *Options) {
			options.Development = false
		})
		assert.Equal(t, logging.InfoLevel, options.level(""))
		logger := options.newZapLogger("foo")
		logger.Info("abc")
		_ = w.Close()
		scanner := bufio.NewScanner(r)
		assert.True(t, scanner.Scan())
		assert.Nil(t, scanner.Err())
		assert.NotEmpty(t, scanner.Text())
		t.Log(scanner.Text())
		m := map[string]any{}
		assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
		assert.NotEmpty(t, m["logger"])
	})
	t.Run("dev", func(t *testing.T) {
		stdout := os.Stdout
		defer func() {
			os.Stdout = stdout
		}()
		r, w, err := os.Pipe()
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = r.Close()
		}()
		os.Stdout = w
		options := NewOptions(func(options *Options) {
			options.Development = true
		})
		assert.Equal(t, logging.InfoLevel, options.level(""))
		logger := options.newZapLogger("foo")
		logger.Info("abc")
		_ = w.Close()
		scanner := bufio.NewScanner(r)
		assert.True(t, scanner.Scan())
		assert.Nil(t, scanner.Err())
		assert.NotEmpty(t, scanner.Text())
		t.Log(scanner.Text())
		m := map[string]any{}
		assert.NotNil(t, json.Unmarshal(scanner.Bytes(), &m))
		assert.Contains(t, scanner.Text(), "foo")
	})
	t.Run("emptyLoggerName", func(t *testing.T) {
		stdout := os.Stdout
		defer func() {
			os.Stdout = stdout
		}()
		r, w, err := os.Pipe()
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = r.Close()
		}()
		os.Stdout = w
		options := NewOptions(func(options *Options) {
			options.Development = false
		})
		assert.Equal(t, logging.InfoLevel, options.level(""))
		logger := options.newZapLogger("")
		logger.Info("abc")
		_ = w.Close()
		scanner := bufio.NewScanner(r)
		assert.True(t, scanner.Scan())
		assert.Nil(t, scanner.Err())
		assert.NotEmpty(t, scanner.Text())
		t.Log(scanner.Text())
		m := map[string]any{}
		assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
		assert.Empty(t, m["logger"])
	})
	t.Run("disableLogger", func(t *testing.T) {
		stdout := os.Stdout
		defer func() {
			os.Stdout = stdout
		}()
		r, w, err := os.Pipe()
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = r.Close()
		}()
		os.Stdout = w
		options := NewOptions(func(options *Options) {
			options.Development = false
			options.DisableLogger = true
		})
		assert.Equal(t, logging.InfoLevel, options.level(""))
		logger := options.newZapLogger("foo")
		logger.Info("abc")
		_ = w.Close()
		scanner := bufio.NewScanner(r)
		assert.True(t, scanner.Scan())
		assert.Nil(t, scanner.Err())
		assert.NotEmpty(t, scanner.Text())
		t.Log(scanner.Text())
		m := map[string]any{}
		assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
		assert.Empty(t, m["logger"])
	})
}
