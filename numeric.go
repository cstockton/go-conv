package conv

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var (
	mathMaxInt  int64
	mathMinInt  int64
	mathMaxUint uint64
	mathIntSize = strconv.IntSize
)

func initIntSizes(size int) {
	switch size {
	case 64:
		mathMaxInt = math.MaxInt64
		mathMinInt = math.MinInt64
		mathMaxUint = math.MaxUint64
	case 32:
		mathMaxInt = math.MaxInt32
		mathMinInt = math.MinInt32
		mathMaxUint = math.MaxUint32
	}
}

func init() {
	// this is so it can be unit tested.
	initIntSizes(mathIntSize)
}

func (c Conv) convStrToFloat64(v string) (float64, bool) {
	if parsed, perr := strconv.ParseFloat(v, 64); perr == nil {
		return parsed, true
	}
	if parsed, perr := c.Bool(v); perr == nil {
		if parsed {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

func (c Conv) convStrToInt64(v string) (int64, error) {
	if parsed, err := strconv.ParseInt(v, 10, 0); err == nil {
		return parsed, nil
	}
	if parsed, err := strconv.ParseFloat(v, 64); err == nil {
		return int64(parsed), nil
	}
	if parsed, err := c.convStrToBool(v); err == nil {
		if parsed {
			return 1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("cannot parse %#v (type string) as integer", v)
}

func (c Conv) convStrToUint64(v string) (uint64, error) {
	if parsed, err := strconv.ParseUint(v, 10, 0); err == nil {
		return parsed, nil
	}
	if parsed, err := strconv.ParseFloat(v, 64); err == nil {
		return uint64(math.Max(0, parsed)), nil
	}
	if parsed, err := c.convStrToBool(v); err == nil {
		if parsed {
			return 1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("cannot parse %#v (type string) as unsigned integer", v)
}

// Float64 attempts to convert the given value to float64, returns the zero
// value and an error on failure.
func (c Conv) Float64(from interface{}) (float64, error) {
	if T, ok := from.(float64); ok {
		return T, nil
	}
	if c, ok := from.(interface {
		Float64() (float64, error)
	}); ok {
		return c.Float64()
	}

	value := reflect.ValueOf(indirect(from))
	kind := value.Kind()
	switch {
	case reflect.String == kind:
		if parsed, ok := c.convStrToFloat64(value.String()); ok {
			return parsed, nil
		}
	case isKindInt(kind):
		return float64(value.Int()), nil
	case isKindUint(kind):
		return float64(value.Uint()), nil
	case isKindFloat(kind):
		return value.Float(), nil
	case isKindComplex(kind):
		return real(value.Complex()), nil
	case reflect.Bool == kind:
		if value.Bool() {
			return 1, nil
		}
		return 0, nil
	case isKindLength(kind):
		return float64(value.Len()), nil
	}
	return 0, newConvErr(from, "float64")
}

// Float32 attempts to convert the given value to Float32, returns the zero
// value and an error on failure.
func (c Conv) Float32(from interface{}) (float32, error) {
	if T, ok := from.(float32); ok {
		return T, nil
	}

	if res, err := c.Float64(from); err == nil {
		if res > math.MaxFloat32 {
			res = math.MaxFloat32
		} else if res < -math.MaxFloat32 {
			res = -math.MaxFloat32
		}
		return float32(res), err
	}
	return 0, newConvErr(from, "float32")
}

// Int64 attempts to convert the given value to int64, returns the zero value
// and an error on failure.
func (c Conv) Int64(from interface{}) (int64, error) {
	if T, ok := from.(string); ok {
		return c.convStrToInt64(T)
	} else if T, ok := from.(int64); ok {
		return T, nil
	}
	if c, ok := from.(interface {
		Int64() (int64, error)
	}); ok {
		return c.Int64()
	}

	value := reflect.ValueOf(indirect(from))
	kind := value.Kind()
	switch {
	case reflect.String == kind:
		return c.convStrToInt64(value.String())
	case isKindInt(kind):
		return value.Int(), nil
	case isKindUint(kind):
		val := value.Uint()
		if val > math.MaxInt64 {
			val = math.MaxInt64
		}
		return int64(val), nil
	case isKindFloat(kind):
		return int64(value.Float()), nil
	case isKindComplex(kind):
		return int64(real(value.Complex())), nil
	case reflect.Bool == kind:
		if value.Bool() {
			return 1, nil
		}
		return 0, nil
	case isKindLength(kind):
		return int64(value.Len()), nil
	}
	return 0, newConvErr(from, "int64")
}

// Int attempts to convert the given value to int, returns the zero value and an
// error on failure.
func (c Conv) Int(from interface{}) (int, error) {
	if T, ok := from.(int); ok {
		return T, nil
	}

	to64, err := c.Int64(from)
	if err != nil {
		return 0, newConvErr(from, "int")
	}
	if to64 > mathMaxInt {
		to64 = mathMaxInt // only possible on 32bit arch
	} else if to64 < mathMinInt {
		to64 = mathMinInt // only possible on 32bit arch
	}
	return int(to64), nil
}

// Int8 attempts to convert the given value to int8, returns the zero value and
// an error on failure.
func (c Conv) Int8(from interface{}) (int8, error) {
	if T, ok := from.(int8); ok {
		return T, nil
	}

	to64, err := c.Int64(from)
	if err != nil {
		return 0, newConvErr(from, "int8")
	}
	if to64 > math.MaxInt8 {
		to64 = math.MaxInt8
	} else if to64 < math.MinInt8 {
		to64 = math.MinInt8
	}
	return int8(to64), nil
}

// Int16 attempts to convert the given value to int16, returns the zero value
// and an error on failure.
func (c Conv) Int16(from interface{}) (int16, error) {
	if T, ok := from.(int16); ok {
		return T, nil
	}

	to64, err := c.Int64(from)
	if err != nil {
		return 0, newConvErr(from, "int16")
	}
	if to64 > math.MaxInt16 {
		to64 = math.MaxInt16
	} else if to64 < math.MinInt16 {
		to64 = math.MinInt16
	}
	return int16(to64), nil
}

// Int32 attempts to convert the given value to int32, returns the zero value
// and an error on failure.
func (c Conv) Int32(from interface{}) (int32, error) {
	if T, ok := from.(int32); ok {
		return T, nil
	}

	to64, err := c.Int64(from)
	if err != nil {
		return 0, newConvErr(from, "int32")
	}
	if to64 > math.MaxInt32 {
		to64 = math.MaxInt32
	} else if to64 < math.MinInt32 {
		to64 = math.MinInt32
	}
	return int32(to64), nil
}

// Uint64 attempts to convert the given value to uint64, returns the zero value
// and an error on failure.
func (c Conv) Uint64(from interface{}) (uint64, error) {
	if T, ok := from.(string); ok {
		return c.convStrToUint64(T)
	} else if T, ok := from.(uint64); ok {
		return T, nil
	}
	if c, ok := from.(interface {
		Uint64() (uint64, error)
	}); ok {
		return c.Uint64()
	}

	value := reflect.ValueOf(indirect(from))
	kind := value.Kind()
	switch {
	case reflect.String == kind:
		return c.convStrToUint64(value.String())
	case isKindUint(kind):
		return value.Uint(), nil
	case isKindInt(kind):
		val := value.Int()
		if val < 0 {
			val = 0
		}
		return uint64(val), nil
	case isKindFloat(kind):
		return uint64(math.Max(0, value.Float())), nil
	case isKindComplex(kind):
		return uint64(math.Max(0, real(value.Complex()))), nil
	case reflect.Bool == kind:
		if value.Bool() {
			return 1, nil
		}
		return 0, nil
	case isKindLength(kind):
		return uint64(value.Len()), nil
	}

	return 0, newConvErr(from, "uint64")
}

// Uint attempts to convert the given value to uint, returns the zero value and
// an error on failure.
func (c Conv) Uint(from interface{}) (uint, error) {
	if T, ok := from.(uint); ok {
		return T, nil
	}

	to64, err := c.Uint64(from)
	if err != nil {
		return 0, newConvErr(from, "uint")
	}
	if to64 > mathMaxUint {
		to64 = mathMaxUint // only possible on 32bit arch
	}
	return uint(to64), nil
}

// Uint8 attempts to convert the given value to uint8, returns the zero value
// and an error on failure.
func (c Conv) Uint8(from interface{}) (uint8, error) {
	if T, ok := from.(uint8); ok {
		return T, nil
	}

	to64, err := c.Uint64(from)
	if err != nil {
		return 0, newConvErr(from, "uint8")
	}
	if to64 > math.MaxUint8 {
		to64 = math.MaxUint8
	}
	return uint8(to64), nil
}

// Uint16 attempts to convert the given value to uint16, returns the zero value
// and an error on failure.
func (c Conv) Uint16(from interface{}) (uint16, error) {
	if T, ok := from.(uint16); ok {
		return T, nil
	}

	to64, err := c.Uint64(from)
	if err != nil {
		return 0, newConvErr(from, "uint16")
	}
	if to64 > math.MaxUint16 {
		to64 = math.MaxUint16
	}
	return uint16(to64), nil
}

// Uint32 attempts to convert the given value to uint32, returns the zero value
// and an error on failure.
func (c Conv) Uint32(from interface{}) (uint32, error) {
	if T, ok := from.(uint32); ok {
		return T, nil
	}

	to64, err := c.Uint64(from)
	if err != nil {
		return 0, newConvErr(from, "uint32")
	}
	if to64 > math.MaxUint32 {
		to64 = math.MaxUint32
	}
	return uint32(to64), nil
}
