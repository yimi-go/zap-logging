package zap

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/yimi-go/logging"
)

func mapZapField(field logging.Field) zapcore.Field {
	switch field.Type() {
	case logging.BinaryType:
		return zap.Binary(field.Key(), field.Value().([]byte))
	case logging.BoolType:
		return zap.Bool(field.Key(), field.Value().(bool))
	case logging.Complex128Type:
		return zap.Complex128(field.Key(), field.Value().(complex128))
	case logging.Complex64Type:
		return zap.Complex64(field.Key(), field.Value().(complex64))
	case logging.DurationType:
		return zap.Duration(field.Key(), field.Value().(time.Duration))
	case logging.Float64Type:
		return zap.Float64(field.Key(), field.Value().(float64))
	case logging.Float32Type:
		return zap.Float32(field.Key(), field.Value().(float32))
	case logging.Int64Type:
		return zap.Int64(field.Key(), field.Value().(int64))
	case logging.Int32Type:
		return zap.Int32(field.Key(), field.Value().(int32))
	case logging.Int16Type:
		return zap.Int16(field.Key(), field.Value().(int16))
	case logging.Int8Type:
		return zap.Int8(field.Key(), field.Value().(int8))
	case logging.StringType:
		return zap.String(field.Key(), field.Value().(string))
	case logging.TimeType:
		return zap.Time(field.Key(), field.Value().(time.Time))
	case logging.Uint64Type:
		return zap.Uint64(field.Key(), field.Value().(uint64))
	case logging.Uint32Type:
		return zap.Uint32(field.Key(), field.Value().(uint32))
	case logging.Uint16Type:
		return zap.Uint16(field.Key(), field.Value().(uint16))
	case logging.Uint8Type:
		return zap.Uint8(field.Key(), field.Value().(uint8))
	case logging.UintptrType:
		return zap.Uintptr(field.Key(), field.Value().(uintptr))
	case logging.StringerType:
		return zap.Stringer(field.Key(), field.Value().(fmt.Stringer))
	case logging.ErrorType:
		return zap.NamedError(field.Key(), field.Value().(error))
	case logging.StackType:
		return zap.StackSkip(field.Key(), field.Value().(int)+3)
	default:
		return zap.Any(field.Key(), field.Value())
	}
}
