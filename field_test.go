package zap

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/yimi-go/logging"
)

type hanBool bool

func (h hanBool) String() string {
	if h {
		return "真"
	}
	return "假"
}

type jsonLine struct {
	Stack string `json:"stack,omitempty"`
}

func Test_mapZapField(t *testing.T) {
	now := time.Now()
	up := uintptr(unsafe.Pointer(&now))
	type args struct {
		field logging.Field
	}
	tests := []struct {
		name string
		args args
		want zapcore.Field
	}{
		{
			name: "binary",
			args: args{logging.Binary("key", []byte("abc"))},
			want: zap.Binary("key", []byte("abc")),
		},
		{
			name: "bool",
			args: args{logging.Bool("key", true)},
			want: zap.Bool("key", true),
		},
		{
			name: "c128",
			args: args{logging.Complex128("key", 1+2i)},
			want: zap.Complex128("key", 1+2i),
		},
		{
			name: "c64",
			args: args{logging.Complex64("key", (complex64)(1+2i))},
			want: zap.Complex64("key", (complex64)(1+2i)),
		},
		{
			name: "duration",
			args: args{logging.Duration("key", time.Second)},
			want: zap.Duration("key", time.Second),
		},
		{
			name: "f64",
			args: args{logging.Float64("key", 1.2)},
			want: zap.Float64("key", 1.2),
		},
		{
			name: "f32",
			args: args{logging.Float32("key", float32(1.2))},
			want: zap.Float32("key", float32(1.2)),
		},
		{
			name: "i64",
			args: args{logging.Int64("key", int64(2))},
			want: zap.Int64("key", int64(2)),
		},
		{
			name: "i32",
			args: args{logging.Int32("key", int32(2))},
			want: zap.Int32("key", int32(2)),
		},
		{
			name: "i16",
			args: args{logging.Int16("key", int16(2))},
			want: zap.Int16("key", int16(2)),
		},
		{
			name: "i8",
			args: args{logging.Int8("key", int8(2))},
			want: zap.Int8("key", int8(2)),
		},
		{
			name: "string",
			args: args{logging.String("key", "val")},
			want: zap.String("key", "val"),
		},
		{
			name: "time",
			args: args{logging.Time("key", now)},
			want: zap.Time("key", now),
		},
		{
			name: "u64",
			args: args{logging.Uint64("key", uint64(2))},
			want: zap.Uint64("key", uint64(2)),
		},
		{
			name: "u32",
			args: args{logging.Uint32("key", uint32(2))},
			want: zap.Uint32("key", uint32(2)),
		},
		{
			name: "u16",
			args: args{logging.Uint16("key", uint16(2))},
			want: zap.Uint16("key", uint16(2)),
		},
		{
			name: "u8",
			args: args{logging.Uint8("key", uint8(2))},
			want: zap.Uint8("key", uint8(2)),
		},
		{
			name: "uintptr",
			args: args{logging.Uintptr("key", up)},
			want: zap.Uintptr("key", up),
		},
		{
			name: "stringer",
			args: args{logging.Stringer("key", hanBool(true))},
			want: zap.Stringer("key", hanBool(true)),
		},
		{
			name: "error",
			args: args{logging.Error(io.EOF)},
			want: zap.Error(io.EOF),
		},
		{
			name: "map",
			args: args{logging.Any("key", map[string]any{"a": 1})},
			want: zap.Any("key", map[string]any{"a": 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mapZapField(tt.args.field), "mapZapField(%v)", tt.args.field)
		})
	}
}

func Test_mapZapField_stack(t *testing.T) {
	origin := os.Stdout
	defer func() {
		os.Stdout = origin
	}()
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	factory := NewFactory(NewOptions())
	logger := factory.Logger("foo")
	logger.Infow("hello", logging.Stack("stack"))
	scanner := bufio.NewScanner(r)
	assert.True(t, scanner.Scan())
	assert.Nil(t, scanner.Err())
	text := scanner.Text()
	t.Log(text)
	jl := jsonLine{}
	err = json.Unmarshal(scanner.Bytes(), &jl)
	assert.Nil(t, err)
	assert.NotEmpty(t, jl.Stack)
	t.Log(jl.Stack)
	split := strings.Split(jl.Stack, "\n")
	assert.Contains(t, split[0], "Test_mapZapField_stack")
}
