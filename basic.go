package conv

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"time"
)

func (c Conv) convStrToBool(v string) (bool, error) {
	// @TODO Need to find a clean way to expose the truth list to be modified by
	// API to allow INTL.
	if 1 > len(v) || len(v) > 5 {
		return false, fmt.Errorf("cannot parse type string with len %d as bool", len(v))
	}

	// @TODO lut
	switch v {
	case "1", "t", "T", "true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES":
		return true, nil
	case "0", "f", "F", "false", "False", "FALSE", "n", "N", "no", "No", "NO":
		return false, nil
	}
	return false, fmt.Errorf("cannot parse %#v (type string) as bool", v)
}

func (c Conv) convNumToBool(k reflect.Kind, value reflect.Value) (bool, bool) {
	switch {
	case isKindInt(k):
		return 0 != value.Int(), true
	case isKindUint(k):
		return 0 != value.Uint(), true
	case isKindFloat(k):
		T := value.Float()
		if math.IsNaN(T) {
			return false, true
		}
		return 0 != T, true
	case isKindComplex(k):
		T := value.Complex()
		if cmplx.IsNaN(T) {
			return false, true
		}
		return 0 != real(T), true
	}
	return false, false
}

// Bool attempts to convert the given value to bool, returns the zero value
// and an error on failure.
func (c Conv) Bool(from interface{}) (bool, error) {
	if T, ok := from.(string); ok {
		return c.convStrToBool(T)
	} else if T, ok := from.(bool); ok {
		return T, nil
	} else if c, ok := from.(interface {
		Bool() (bool, error)
	}); ok {
		return c.Bool()
	}

	value := reflect.ValueOf(indirect(from))
	kind := value.Kind()
	switch {
	case reflect.String == kind:
		return c.convStrToBool(value.String())
	case isKindNumeric(kind):
		if parsed, ok := c.convNumToBool(kind, value); ok {
			return parsed, nil
		}
	case reflect.Bool == kind:
		return value.Bool(), nil
	case isKindLength(kind):
		return value.Len() > 0, nil
	case reflect.Struct == kind && value.CanInterface():
		v := value.Interface()
		if t, ok := v.(time.Time); ok {
			return emptyTime != t, nil
		}
	}
	return false, newConvErr(from, "bool")
}

// String returns the string representation from the given interface{} value
// and can not currently fail. Although an error is currently provided only for
// API cohesion you should still check it to be future proof.
func (c Conv) String(from interface{}) (string, error) {
	if T, ok := from.(string); ok {
		return T, nil
	} else if T, ok := from.([]byte); ok {
		return string(T), nil
	} else if c, ok := from.(interface {
		// @TODO This aligns with the API, but not with Go interfaces.
		String() (string, error)
	}); ok {
		return c.String()
	}

	switch T := from.(type) {
	case *[]byte:
		// @TODO Maybe validate the bytes are valid runes
		return string(*T), nil
	case *string:
		return *T, nil
	}
	return fmt.Sprintf("%v", from), nil
}
