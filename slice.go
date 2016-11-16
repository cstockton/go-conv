package conv

import (
	"fmt"

	iter "github.com/cstockton/go-iter"
)

// Slice will perform conversion by inferring the element type from the given
// slice and taking values from the given interface.
func (c Conv) Slice(into interface{}, from interface{}) error {
	return recoverFn(func() error {
		fn := sliceIterFn(c, into)
		if fn == nil {
			return fmt.Errorf("%T is not a pointer to a supported slice type", into)
		}
		return iter.Walk(from, fn)
	})
}
