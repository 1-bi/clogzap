package log_zap

import (
	"go.uber.org/zap"
	"time"
)

type zapLoggerBean struct {

	// define name array , value is field type
	fieldProps map[string]zap.Field
}

func (myself *zapLoggerBean) convertToFields() []zap.Field {

	var fields = make([]zap.Field, 0)
	for _, vf := range myself.fieldProps {
		fields = append(fields, vf)
	}

	return fields
}

func (myself *zapLoggerBean) LogBinary(key string, value []byte) {
	myself.fieldProps[key] = zap.Binary(key, value)
}

func (myself *zapLoggerBean) LogByteString(key string, values []byte) {
	myself.fieldProps[key] = zap.ByteString(key, values)
}

func (myself *zapLoggerBean) LogByteStringArray(key string, values [][]byte) {
	myself.fieldProps[key] = zap.ByteStrings(key, values)
}

func (myself *zapLoggerBean) LogString(key string, value string) {
	myself.fieldProps[key] = zap.String(key, value)
}

func (myself *zapLoggerBean) LogStringArray(key string, values []string) {
	myself.fieldProps[key] = zap.Strings(key, values)
}

func (myself *zapLoggerBean) LogBool(key string, value bool) {
	myself.fieldProps[key] = zap.Bool(key, value)
}

func (myself *zapLoggerBean) LogBoolArray(key string, values []bool) {
	myself.fieldProps[key] = zap.Bools(key, values)
}

// --- int

func (myself *zapLoggerBean) LogInt8(key string, value int8) {
	myself.fieldProps[key] = zap.Int8(key, value)
}

func (myself *zapLoggerBean) LogInt8Array(key string, nums []int8) {
	myself.fieldProps[key] = zap.Int8s(key, nums)
}

func (myself *zapLoggerBean) LogInt(key string, value int) {
	myself.fieldProps[key] = zap.Int(key, value)
}

func (myself *zapLoggerBean) LogIntArray(key string, nums []int) {
	myself.fieldProps[key] = zap.Ints(key, nums)
}

func (myself *zapLoggerBean) LogInt16(key string, value int16) {
	myself.fieldProps[key] = zap.Int16(key, value)
}

func (myself *zapLoggerBean) LogInt16Array(key string, nums []int16) {
	myself.fieldProps[key] = zap.Int16s(key, nums)
}

func (myself *zapLoggerBean) LogInt32(key string, value int32) {
	myself.fieldProps[key] = zap.Int32(key, value)
}

func (myself *zapLoggerBean) LogInt32Array(key string, nums []int32) {
	myself.fieldProps[key] = zap.Int32s(key, nums)
}

func (myself *zapLoggerBean) LogInt64(key string, value int64) {
	myself.fieldProps[key] = zap.Int64(key, value)
}

func (myself *zapLoggerBean) LogInt64Array(key string, nums []int64) {
	myself.fieldProps[key] = zap.Int64s(key, nums)
}

// --- uint
func (myself *zapLoggerBean) LogUint8(key string, value uint8) {
	myself.fieldProps[key] = zap.Uint8(key, value)
}

func (myself *zapLoggerBean) LogUint8Array(key string, nums []uint8) {
	myself.fieldProps[key] = zap.Uint8s(key, nums)
}

func (myself *zapLoggerBean) LogUint(key string, value uint) {
	myself.fieldProps[key] = zap.Uint(key, value)
}

func (myself *zapLoggerBean) LogUintArray(key string, nums []uint) {
	myself.fieldProps[key] = zap.Uints(key, nums)
}

func (myself *zapLoggerBean) LogUint16(key string, value uint16) {
	myself.fieldProps[key] = zap.Uint16(key, value)
}

func (myself *zapLoggerBean) LogUint16Array(key string, nums []uint16) {
	myself.fieldProps[key] = zap.Uint16s(key, nums)
}

func (myself *zapLoggerBean) LogUint32(key string, value uint32) {
	myself.fieldProps[key] = zap.Uint32(key, value)
}

func (myself *zapLoggerBean) LogUint32Array(key string, nums []uint32) {
	myself.fieldProps[key] = zap.Uint32s(key, nums)
}

// --- float
func (myself *zapLoggerBean) LogFloat32(key string, value float32) {
	myself.fieldProps[key] = zap.Float32(key, value)
}

func (myself *zapLoggerBean) LogFloat32Array(key string, values []float32) {
	myself.fieldProps[key] = zap.Float32s(key, values)
}

func (myself *zapLoggerBean) LogFloat64(key string, value float64) {
	myself.fieldProps[key] = zap.Float64(key, value)
}

func (myself *zapLoggerBean) LogFloat64Array(key string, values []float64) {
	myself.fieldProps[key] = zap.Float64s(key, values)
}

// --- time or duration

func (myself *zapLoggerBean) LogDuration(key string, value time.Duration) {
	myself.fieldProps[key] = zap.Duration(key, value)
}

func (myself *zapLoggerBean) LogDurationArray(key string, values []time.Duration) {
	myself.fieldProps[key] = zap.Durations(key, values)
}

func (myself *zapLoggerBean) LogTime(key string, value time.Time) {
	myself.fieldProps[key] = zap.Time(key, value)
}

func (myself *zapLoggerBean) LogTimeArray(key string, values []time.Time) {
	myself.fieldProps[key] = zap.Times(key, values)
}
