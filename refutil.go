package conv

import (
	"fmt"
	"reflect"
)

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
