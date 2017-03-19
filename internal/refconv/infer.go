package refconv

import (
	"fmt"
	"reflect"
)

// Infer will perform conversion by inferring the conversion operation from
// the T of `into`.
func (c Conv) Infer(into, from interface{}) error {
	var (
		v   interface{}
		err error
	)

	value := reflect.ValueOf(into)
	if !value.IsValid() {
		return fmt.Errorf("%T is not a valid value", into)
	}

	// Check for a map before dereference
	kind := value.Kind()
	if kind != reflect.Ptr {
		return fmt.Errorf(`cannot infer conversion for non-pointer %v (type %[1]T)`, into)
	}

	value = value.Elem()
	kind = value.Kind()
	switch kind {
	case reflect.Bool:
		v, err = c.Bool(from)
	case reflect.Float32:
		v, err = c.Float32(from)
	case reflect.Float64:
		v, err = c.Float64(from)
	case reflect.Int:
		v, err = c.Int(from)
	case reflect.Int8:
		v, err = c.Int8(from)
	case reflect.Int16:
		v, err = c.Int16(from)
	case reflect.Int32:
		v, err = c.Int32(from)
	case reflect.String:
		v, err = c.String(from)
	case reflect.Uint:
		v, err = c.Uint(from)
	case reflect.Uint8:
		v, err = c.Uint8(from)
	case reflect.Uint16:
		v, err = c.Uint16(from)
	case reflect.Uint32:
		v, err = c.Uint32(from)
	case reflect.Uint64:
		v, err = c.Uint64(from)

	// Special cases
	case reflect.Int64:
		if value.Type() == typeOfDuration {
			v, err = c.Duration(from)
		} else {
			v, err = c.Int64(from)
		}
	case reflect.Struct:
		if value.Type() == typeOfTime {
			v, err = c.Time(from)
		}
	default:
		err = fmt.Errorf(`cannot infer conversion for %v (type %[1]T)`, into)
	}
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(v))
	return nil
}
