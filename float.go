package conv

import (
	"math"
	"reflect"
	"strconv"
)

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
