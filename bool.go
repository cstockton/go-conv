package conv

import "time"

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
		switch T {
		case "1", "t", "T", "true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES":
			return true
		case "0", "f", "F", "false", "False", "FALSE", "n", "N", "no", "No", "NO":
			return false
		}
	}
	return Int64(value) != 0
}
