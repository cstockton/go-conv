package conv

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

var (
	nilValues = []interface{}{
		(*interface{})(nil), (**interface{})(nil), (***interface{})(nil),
		(func())(nil), (*func())(nil), (**func())(nil), (***func())(nil),
		(chan int)(nil), (*chan int)(nil), (**chan int)(nil), (***chan int)(nil),
		([]int)(nil), (*[]int)(nil), (**[]int)(nil), (***[]int)(nil),
		(map[int]int)(nil), (*map[int]int)(nil), (**map[int]int)(nil),
		(***map[int]int)(nil),
	}
)

func TestIndirect(t *testing.T) {
	type testIndirectCircular *testIndirectCircular
	teq := func(t testing.TB, exp, got interface{}) {
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("DeepEqual failed:\n  exp: %#v\n  got: %#v", exp, got)
		}
	}

	t.Run("Basic", func(t *testing.T) {
		int64v := int64(123)
		int64vp := &int64v
		int64vpp := &int64vp
		int64vppp := &int64vpp
		int64vpppp := &int64vppp
		teq(t, indirect(int64v), int64v)
		teq(t, indirect(int64vp), int64v)
		teq(t, indirect(int64vpp), int64v)
		teq(t, indirect(int64vppp), int64v)
		teq(t, indirect(int64vpppp), int64v)
	})
	t.Run("Nils", func(t *testing.T) {
		for _, n := range nilValues {
			indirect(n)
		}
	})
	t.Run("Circular", func(t *testing.T) {
		var circular testIndirectCircular
		circular = &circular
		teq(t, indirect(circular), circular)
	})
}

func TestRecoverFn(t *testing.T) {
	t.Run("CallsFunc", func(t *testing.T) {
		var called bool

		err := recoverFn(func() error {
			called = true
			return nil
		})
		if err != nil {
			t.Error("expected no error in recoverFn()")
		}
		if !called {
			t.Error("Expected recoverFn() to call func")
		}
	})
	t.Run("PropagatesError", func(t *testing.T) {
		err := fmt.Errorf("expect this error")
		rerr := recoverFn(func() error {
			return err
		})
		if err != rerr {
			t.Error("expected recoverFn() to propagate")
		}
	})
	t.Run("PropagatesPanicError", func(t *testing.T) {
		err := fmt.Errorf("expect this error")
		rerr := recoverFn(func() error {
			panic(err)
		})
		if err != rerr {
			t.Error("Expected recoverFn() to propagate")
		}
	})
	t.Run("PropagatesRuntimeError", func(t *testing.T) {
		err := recoverFn(func() error {
			sl := []int{}
			_ = sl[0]
			return nil
		})
		if err == nil {
			t.Error("expected runtime error to propagate")
		}
		if _, ok := err.(runtime.Error); !ok {
			t.Error("expected runtime error to retain type type")
		}
	})
	t.Run("PropagatesString", func(t *testing.T) {
		exp := "panic: string type panic"
		rerr := recoverFn(func() error {
			panic("string type panic")
		})
		if exp != rerr.Error() {
			t.Errorf("expected recoverFn() to return %v, got: %v", exp, rerr)
		}
	})
}

func TestKind(t *testing.T) {
	var (
		intKinds = []reflect.Kind{
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
		uintKinds = []reflect.Kind{
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64}
		floatKinds   = []reflect.Kind{reflect.Float32, reflect.Float64}
		complexKinds = []reflect.Kind{reflect.Complex64, reflect.Complex128}
		lengthKinds  = []reflect.Kind{
			reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String}
		nilKinds = []reflect.Kind{
			reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice,
			reflect.Ptr}
	)
	type kindFunc func(k reflect.Kind) bool
	type testKind struct {
		exp   bool
		kinds []reflect.Kind
		f     []kindFunc
	}
	tnew := func(exp bool, k []reflect.Kind, f ...kindFunc) testKind {
		return testKind{exp, k, f}
	}
	tests := []testKind{
		tnew(true, intKinds, isKindInt, isKindNumeric),
		tnew(false, intKinds, isKindComplex, isKindFloat, isKindUint, isKindLength, isKindNil),
		tnew(true, uintKinds, isKindNumeric, isKindUint),
		tnew(false, uintKinds, isKindComplex, isKindFloat, isKindInt, isKindLength, isKindNil),
		tnew(true, floatKinds, isKindFloat, isKindNumeric),
		tnew(false, floatKinds, isKindComplex, isKindInt, isKindUint, isKindLength, isKindNil),
		tnew(true, complexKinds, isKindComplex, isKindNumeric),
		tnew(false, complexKinds, isKindFloat, isKindInt, isKindUint, isKindLength, isKindNil),
		tnew(true, lengthKinds, isKindLength),
		tnew(true, nilKinds, isKindNil),
	}
	for _, tc := range tests {
		for _, f := range tc.f {
			t.Run(fmt.Sprintf("%v", tc.kinds[0]), func(t *testing.T) {
				for _, kind := range tc.kinds {
					if got := f(kind); got != tc.exp {
						t.Errorf("%#v(%v)\nexp: %v\ngot: %v", f, kind, tc.exp, got)
					}
				}
			})
		}
	}
}
