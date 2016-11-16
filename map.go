package conv

import (
	"fmt"

	iter "github.com/cstockton/go-iter"
)

// Map will perform conversion by inferring the key and element types from the
// given map and taking values from the given interface.
func (c Conv) Map(into, from interface{}) error {
	return recoverFn(func() error {
		fn := mapIterFn(c, into)
		if fn == nil {
			return fmt.Errorf("%T is not a supported map type", into)
		}
		return iter.Walk(from, fn)
	})
}
