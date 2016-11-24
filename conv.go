// Package conv provides fast and intuitive conversions across Go types.
package conv

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"time"
)

// Conv implements the Converter interface. It does not require initialization
// or share state and is safe for use by multiple Goroutines.
type Conv struct{}

var (

	// DefaultConv is used by the top level functions in this package. The callers
	// will discard any errors.
	DefaultConv Converter = Conv{}
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

// Converter supports conversion to basic types, that is Boolean, Numeric and
// Strings. As a special case it may convert to the time.Time structure. It is
// the primary user facing interface for this library.
type Converter interface {

	// Map will perform conversion by inferring the key and element types from the
	// given map and taking values from the given interface.
	Map(into, from interface{}) error

	// Slice will perform conversion by inferring the element type from the given
	// slice and taking values from the given interface.
	Slice(into, from interface{}) error

	// Bool returns the bool representation from the given interface value.
	// Returns the default value of false and an error on failure.
	Bool(from interface{}) (to bool, err error)

	// Duration returns the time.Duration representation from the given
	// interface{} value. Returns the default value of 0 and an error on failure.
	Duration(from interface{}) (to time.Duration, err error)

	// String returns the string representation from the given interface
	// value and can not fail. An error is provided only for API cohesion.
	String(from interface{}) (to string, err error)

	// Time returns the time.Time{} representation from the given interface
	// value. Returns an empty time.Time struct and an error on failure.
	Time(from interface{}) (to time.Time, err error)

	// Float32 returns the float32 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Float32(from interface{}) (to float32, err error)

	// Float64 returns the float64 representation from the given interface
	// value. Returns the default value of 0 and an error on failure.
	Float64(from interface{}) (to float64, err error)

	// Int returns the int representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Int(from interface{}) (to int, err error)

	// Int8 returns the int8 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Int8(from interface{}) (to int8, err error)

	// Int16 returns the int16 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Int16(from interface{}) (to int16, err error)

	// Int32 returns the int32 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Int32(from interface{}) (to int32, err error)

	// Int64 returns the int64 representation from the given interface
	// value. Returns the default value of 0 and an error on failure.
	Int64(from interface{}) (to int64, err error)

	// Uint returns the uint representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Uint(from interface{}) (to uint, err error)

	// Uint8 returns the uint8 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Uint8(from interface{}) (to uint8, err error)

	// Uint16 returns the uint16 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Uint16(from interface{}) (to uint16, err error)

	// Uint32 returns the uint32 representation from the given empty interface
	// value. Returns the default value of 0 and an error on failure.
	Uint32(from interface{}) (to uint32, err error)

	// Uint64 returns the uint64 representation from the given interface
	// value. Returns the default value of 0 and an error on failure.
	Uint64(from interface{}) (to uint64, err error)
}

// Slice will perform conversion by inferring the element type from the given
// slice. The from element is traversed recursively and the behavior if the
// value is mutated during iteration is undefined, though at worst an error
// will be returned as this library will never panic.
//
// An error is returned if the below restrictions are not met:
//
//   - It must be a pointer to a slice, it does not have to be initialized
//   - The element must be a T or *T of a type supported by this library
//
// Example:
//
//   var into []int64
//   err := conv.Slice(&into, []string{"12", "345", "6789"})
//   // into -> []int64{12, 234, 6789}
//
// See examples for more usages.
func Slice(into, from interface{}) error {
	return DefaultConv.Slice(into, from)
}

// Map will perform conversion by inferring the key and element types from the
// given map. The from element is traversed recursively and the behavior if
// the value is mutated during iteration is undefined, though at worst an
// error will be returned as this library will never panic.
//
// An error is returned if the below restrictions are not met:
//
//   - It must be a non-pointer, non-nil initialized map
//   - Both the key and element T must be supported by this library
//   - The key must be a value T, the element may be a T or *T
//
// Example:
//
//   into := make(map[string]int64)
//   err := conv.Map(into, []string{"12", "345", "6789"})
//   // into -> map[string]int64{"0": 12, "1", 234, "2", 6789}
//
// See examples for more usages.
func Map(into, from interface{}) error {
	return DefaultConv.Map(into, from)
}

// Bool will convert the given value to a bool, returns the default value of
// false if a conversion can not be made.
func Bool(from interface{}) bool {
	to, _ := DefaultConv.Bool(from)
	return to
}

// Duration will convert the given value to a time.Duration, returns the default
// value of 0ns if a conversion can not be made.
func Duration(from interface{}) time.Duration {
	to, _ := DefaultConv.Duration(from)
	return to
}

// String will convert the given value to a string, returns the default value
// of "" if a conversion can not be made.
func String(from interface{}) string {
	to, _ := DefaultConv.String(from)
	return to
}

// Time will convert the given value to a time.Time, returns the empty struct
// time.Time{} if a conversion can not be made.
func Time(from interface{}) time.Time {
	to, _ := DefaultConv.Time(from)
	return to
}

// Float32 will convert the given value to a float32, returns the default value
// of 0.0 if a conversion can not be made.
func Float32(from interface{}) float32 {
	to, _ := DefaultConv.Float32(from)
	return to
}

// Float64 will convert the given value to a float64, returns the default value
// of 0.0 if a conversion can not be made.
func Float64(from interface{}) float64 {
	to, _ := DefaultConv.Float64(from)
	return to
}

// Int will convert the given value to a int, returns the default value of 0 if
// a conversion can not be made.
func Int(from interface{}) int {
	to, _ := DefaultConv.Int(from)
	return to
}

// Int8 will convert the given value to a int8, returns the default value of 0
// if a conversion can not be made.
func Int8(from interface{}) int8 {
	to, _ := DefaultConv.Int8(from)
	return to
}

// Int16 will convert the given value to a int16, returns the default value of 0
// if a conversion can not be made.
func Int16(from interface{}) int16 {
	to, _ := DefaultConv.Int16(from)
	return to
}

// Int32 will convert the given value to a int32, returns the default value of 0
// if a conversion can not be made.
func Int32(from interface{}) int32 {
	to, _ := DefaultConv.Int32(from)
	return to
}

// Int64 will convert the given value to a int64, returns the default value of 0
// if a conversion can not be made.
func Int64(from interface{}) int64 {
	to, _ := DefaultConv.Int64(from)
	return to
}

// Uint will convert the given value to a uint, returns the default value of 0
// if a conversion can not be made.
func Uint(from interface{}) uint {
	to, _ := DefaultConv.Uint(from)
	return to
}

// Uint8 will convert the given value to a uint8, returns the default value of 0
// if a conversion can not be made.
func Uint8(from interface{}) uint8 {
	to, _ := DefaultConv.Uint8(from)
	return to
}

// Uint16 will convert the given value to a uint16, returns the default value of
// 0 if a conversion can not be made.
func Uint16(from interface{}) uint16 {
	to, _ := DefaultConv.Uint16(from)
	return to
}

// Uint32 will convert the given value to a uint32, returns the default value of
// 0 if a conversion can not be made.
func Uint32(from interface{}) uint32 {
	to, _ := DefaultConv.Uint32(from)
	return to
}

// Uint64 will convert the given value to a uint64, returns the default value of
// 0 if a conversion can not be made.
func Uint64(from interface{}) uint64 {
	to, _ := DefaultConv.Uint64(from)
	return to
}

func newConvErr(from interface{}, to string) error {
	return fmt.Errorf("cannot convert %#v (type %[1]T) to %v", from, to)
}

// indirect( will perform recursive indirection on the given value. It should
// never panic and will return a value unless indirection is impossible due to
// infinite recursion in cases like `type Element *Element`.
func indirect(value interface{}) interface{} {

	// Just to be safe, recursion should not be possible but I may be
	// missing an edge case.
	for {

		val := reflect.ValueOf(value)
		if !val.IsValid() || val.Kind() != reflect.Ptr {
			// Value is not a pointer.
			return value
		}

		res := reflect.Indirect(val)
		if !res.IsValid() || !res.CanInterface() {
			// Invalid value or can't be returned as interface{}.
			return value
		}

		// Test for a circular type.
		if res.Kind() == reflect.Ptr && val.Pointer() == res.Pointer() {
			return value
		}

		// Next round.
		value = res.Interface()
	}
}

// recoverFn will attempt to execute f, if f return a non-nil error it will be
// returned. If f panics this function will attempt to recover() and return a
// error instead.
func recoverFn(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch T := r.(type) {
			case error:
				err = T
			default:
				err = fmt.Errorf("panic: %v", r)
			}
		}
	}()
	err = f()
	return
}

// isKindComplex returns true if the given Kind is a complex value.
func isKindComplex(k reflect.Kind) bool {
	return reflect.Complex64 == k || k == reflect.Complex128
}

// isKindFloat returns true if the given Kind is a float value.
func isKindFloat(k reflect.Kind) bool {
	return reflect.Float32 == k || k == reflect.Float64
}

// isKindInt returns true if the given Kind is a int value.
func isKindInt(k reflect.Kind) bool {
	return reflect.Int <= k && k <= reflect.Int64
}

// isKindUint returns true if the given Kind is a uint value.
func isKindUint(k reflect.Kind) bool {
	return reflect.Uint <= k && k <= reflect.Uint64
}

// isKindNumeric returns true if the given Kind is a numeric value.
func isKindNumeric(k reflect.Kind) bool {
	return (reflect.Int <= k && k <= reflect.Uint64) ||
		(reflect.Float32 <= k && k <= reflect.Complex128)
}

// isKindNil will return true if the Kind is a chan, func, interface, map,
// pointer, or slice value, false otherwise.
func isKindNil(k reflect.Kind) bool {
	return (reflect.Chan <= k && k <= reflect.Slice) || k == reflect.UnsafePointer
}

// isKindLength will return true if the Kind has a length.
func isKindLength(k reflect.Kind) bool {
	return reflect.Array == k || reflect.Chan == k || reflect.Map == k ||
		reflect.Slice == k || reflect.String == k
}
