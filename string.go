package conv

import "fmt"

// StringConverter interface allows a value to be converted to a string.
type StringConverter interface {
	String() string
}

// String converts the given value to a string.
func String(value interface{}) string {
	value = Indirect(value)

	switch T := value.(type) {
	case StringConverter:
		if T != nil {
			return T.String()
		}
	case []byte:
		return string(T)
	case string:
		return T
	}
	return fmt.Sprintf("%v", value)
}
