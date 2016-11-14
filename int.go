package conv

import (
	"math"
	"reflect"
	"strconv"
)

func (c Conv) convStrToInt64(v string) (int64, bool) {
	if parsed, perr := strconv.ParseInt(v, 10, 0); perr == nil {
		return parsed, true
	}
	if parsed, perr := strconv.ParseFloat(v, 64); perr == nil {
		return int64(parsed), true
	}
	if parsed, perr := c.Bool(v); perr == nil {
		if parsed {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

// Int64 attempts to convert the given value to int64, returns the zero value
// and an error on failure.
func (c Conv) Int64(from interface{}) (int64, error) {
	if T, ok := from.(int64); ok {
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
		if parsed, ok := c.convStrToInt64(value.String()); ok {
			return parsed, nil
		}
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
