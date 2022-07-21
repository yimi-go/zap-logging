package zap

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yimi-go/logging"
)

func prepareZapFactory(name string, level logging.Level, opts ...Option) (*zapFactory, io.WriteCloser, io.ReadCloser) {
	options := NewOptions(func(options *Options) {
		options.Levels[name] = level
	})
	for _, o := range opts {
		o(options)
	}
	factory := NewFactory(options)
	origin := os.Stdout
	defer func() {
		os.Stdout = origin
	}()
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	_ = factory.zap(name)
	return factory, w, r
}

func Test_zapLogger_Debug(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.DebugLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debug("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "ab", m["msg"])
	assert.Equal(t, "DEBUG", m["level"])
}

func Test_zapLogger_Debug2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debug("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Debugln(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.DebugLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debugln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "a b", m["msg"])
	assert.Equal(t, "DEBUG", m["level"])
}

func Test_zapLogger_Debugln2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debugln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Debugf(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.DebugLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debugf("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "(a, b)", m["msg"])
	assert.Equal(t, "DEBUG", m["level"])
}

func Test_zapLogger_Debugf2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debugf("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Debugw(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.DebugLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debugw("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "hello", m["msg"])
	assert.Equal(t, "bar", m["foo"])
	assert.Equal(t, "DEBUG", m["level"])
}

func Test_zapLogger_Debugw2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Debugw("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Info(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Info("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "ab", m["msg"])
	assert.Equal(t, "INFO", m["level"])
}

func Test_zapLogger_Info2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Info("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Infoln(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Infoln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "a b", m["msg"])
	assert.Equal(t, "INFO", m["level"])
}

func Test_zapLogger_Infoln2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Infoln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Infof(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Infof("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "(a, b)", m["msg"])
	assert.Equal(t, "INFO", m["level"])
}

func Test_zapLogger_Infof2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Infof("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Infow(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Infow("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "hello", m["msg"])
	assert.Equal(t, "bar", m["foo"])
	assert.Equal(t, "INFO", m["level"])
}

func Test_zapLogger_Infow2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Infow("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Warn(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warn("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "ab", m["msg"])
	assert.Equal(t, "WARN", m["level"])
}

func Test_zapLogger_Warn2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warn("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Warnln(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warnln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "a b", m["msg"])
	assert.Equal(t, "WARN", m["level"])
}

func Test_zapLogger_Warnln2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warnln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Warnf(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warnf("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "(a, b)", m["msg"])
	assert.Equal(t, "WARN", m["level"])
}

func Test_zapLogger_Warnf2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warnf("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Warnw(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warnw("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "hello", m["msg"])
	assert.Equal(t, "bar", m["foo"])
	assert.Equal(t, "WARN", m["level"])
}

func Test_zapLogger_Warnw2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Warnw("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Error(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Error("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "ab", m["msg"])
	assert.Equal(t, "ERROR", m["level"])
}

func Test_zapLogger_Error2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.OffLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Error("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Errorln(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Errorln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "a b", m["msg"])
	assert.Equal(t, "ERROR", m["level"])
}

func Test_zapLogger_Errorln2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.OffLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Errorln("a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Errorf(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Errorf("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "(a, b)", m["msg"])
	assert.Equal(t, "ERROR", m["level"])
}

func Test_zapLogger_Errorf2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.OffLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Errorf("(%s, %s)", "a", "b")
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Errorw(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.ErrorLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Errorw("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "hello", m["msg"])
	assert.Equal(t, "bar", m["foo"])
	assert.Equal(t, "ERROR", m["level"])
}

func Test_zapLogger_Errorw2(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.OffLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}
	l.Errorw("hello", logging.String("foo", "bar"))
	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.False(t, scanner.Scan())
}

func Test_zapLogger_Enabled(t *testing.T) {
	factory, closer, readCloser := prepareZapFactory("foo", logging.WarnLevel)
	defer func() {
		_ = closer.Close()
		_ = readCloser.Close()
	}()
	l := factory.Logger("foo")
	assert.False(t, l.Enabled(logging.DebugLevel))
	assert.False(t, l.Enabled(logging.InfoLevel))
	assert.True(t, l.Enabled(logging.WarnLevel))
	assert.True(t, l.Enabled(logging.ErrorLevel))
	assert.False(t, l.Enabled(logging.OffLevel))
	assert.False(t, l.Enabled(logging.Level(99)))
	assert.False(t, l.Enabled(logging.Level(-99)))
}

func Test_zapLogger_WithField(t *testing.T) {
	factory, writeCloser, readCloser := prepareZapFactory("foo", logging.InfoLevel)
	defer func() {
		_ = readCloser.Close()
	}()
	l := &zapLogger{fields: nil, factory: factory, name: "foo"}

	l2 := l.WithField(logging.String("foo", "bar"))
	l2.Info("hello")

	_ = writeCloser.Close()
	scanner := bufio.NewScanner(readCloser)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	assert.NotEmpty(t, scanner.Text())
	t.Log(scanner.Text())
	m := map[string]any{}
	assert.Nil(t, json.Unmarshal(scanner.Bytes(), &m))
	assert.Equal(t, "foo", m["logger"])
	assert.Equal(t, "hello", m["msg"])
	assert.Equal(t, "bar", m["foo"])
	assert.Equal(t, "INFO", m["level"])
}
