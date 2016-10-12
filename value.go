package conv

import (
	"time"
)

// Value has a single field V which stores an interface. It implements Converter
// by passing the wrapped interface to this packages conversion functions.
type Value struct {
	V interface{}
}

// Flatten will recursively merge any Value.V that is of type Value turning
// Value{...Value{"Foo"}} into Value{"Foo"}.
func (v Value) Flatten() Value {
	for {
		if value, ok := v.V.(Value); ok {
			v.V = value.V
			continue
		} else {
			return v
		}
	}
}

// Indirect returns the underlying interface after indirection.
func (v Value) Indirect() interface{} {
	return Indirect(v.V)
}

// Interface returns the underlying interface{} value in V.
func (v Value) Interface() interface{} {
	return v.V
}

// Bool conversion from the underlying interface{} value in V.
func (v Value) Bool() bool {
	return Bool(v.V)
}

// Complex64 conversion from the underlying interface{} value in V.
func (v Value) Complex64() complex64 {
	return Complex64(v.V)
}

// Complex128 conversion from the underlying interface{} value in V.
func (v Value) Complex128() complex128 {
	return Complex128(v.V)
}

// Duration conversion from the underlying interface{} value in V.
func (v Value) Duration() time.Duration {
	return Duration(v.V)
}

// Float32 conversion from the underlying interface{} value in V.
func (v Value) Float32() float32 {
	return Float32(v.V)
}

// Float64 conversion from the underlying interface{} value in V.
func (v Value) Float64() float64 {
	return Float64(v.V)
}

// Int conversion from the underlying interface{} value in V.
func (v Value) Int() int {
	return Int(v.V)
}

// Int8 conversion from the underlying interface{} value in V.
func (v Value) Int8() int8 {
	return Int8(v.V)
}

// Int16 conversion from the underlying interface{} value in V.
func (v Value) Int16() int16 {
	return Int16(v.V)
}

// Int32 conversion from the underlying interface{} value in V.
func (v Value) Int32() int32 {
	return Int32(v.V)
}

// Int64 conversion from the underlying interface{} value in V.
func (v Value) Int64() int64 {
	return Int64(v.V)
}

// String conversion from the underlying interface{} value in V.
func (v Value) String() string {
	return String(v.V)
}

// Time conversion from the underlying interface{} value in V.
func (v Value) Time() time.Time {
	return Time(v.V)
}

// Uint conversion from the underlying interface{} value in V.
func (v Value) Uint() uint {
	return Uint(v.V)
}

// Uint8 conversion from the underlying interface{} value in V.
func (v Value) Uint8() uint8 {
	return Uint8(v.V)
}

// Uint16 conversion from the underlying interface{} value in V.
func (v Value) Uint16() uint16 {
	return Uint16(v.V)
}

// Uint32 conversion from the underlying interface{} value in V.
func (v Value) Uint32() uint32 {
	return Uint32(v.V)
}

// Uint64 conversion from the underlying interface{} value in V.
func (v Value) Uint64() uint64 {
	return Uint64(v.V)
}
