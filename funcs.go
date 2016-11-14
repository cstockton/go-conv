package conv

import "time"

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
