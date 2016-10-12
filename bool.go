package conv

import (
	"strconv"
	"time"
)

// BoolConverter interface allows a value to be converted to a bool.
type BoolConverter interface {
	Bool() bool
}

// Bool converts the given value to a bool.
func Bool(value interface{}) bool {
	value = Indirect(value)

	switch T := value.(type) {
	case BoolConverter:
		if T != nil {
			return T.Bool()
		}
	case time.Duration:
		return time.Duration(0) != T
	case time.Time:
		return time.Time{} != T
	case bool:
		return T
	case string:
		if parsed, err := strconv.ParseBool(T); err == nil {
			return parsed
		}
	}
	return Int64(value) != 0
}
