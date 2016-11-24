package conv

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/doc"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"math"
	"math/cmplx"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"text/template"
	"time"
)

var (
	summary         string
	assertions      Assertions
	assertionsIndex = make(AssertionsIndex)
	toTheHeap       interface{}
)

const (
	TypeKind reflect.Kind = reflect.UnsafePointer + iota
	DurationKind
	TimeKind
	SliceMask reflect.Kind = (1 << 8)
	MapMask                = (1 << 16)
	ChanMask               = (1 << 24)
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

func TestMain(m *testing.M) {
	chkMathIntSize := mathIntSize
	chkMathMaxInt := mathMaxInt
	chkMathMinInt := mathMinInt
	chkMathMaxUint := mathMaxUint
	chkEmptyTime := time.Time{}

	flag.Parse()
	res := m.Run()

	// validate our max (u|)int sizes don't get written to on accident.
	if chkMathIntSize != mathIntSize {
		panic("chkEmptyTime != emptyTime")
	}
	if chkMathMaxInt != mathMaxInt {
		panic("chkMathMaxInt != mathMaxInt")
	}
	if chkMathMinInt != mathMinInt {
		panic("chkMathMinInt != mathMaxInt")
	}
	if chkMathMaxUint != mathMaxUint {
		panic("chkMathMaxUint != mathMaxUint")
	}
	if chkEmptyTime != emptyTime {
		panic("chkEmptyTime != emptyTime")
	}

	os.Exit(res)
}

func TestString(t *testing.T) {
	var c Conv
	t.Run("String", func(t *testing.T) {
		if n := assertions.EachOf(reflect.String, func(a *Assertion, e Expecter) {
			if err := e.Expect(c.String(a.From)); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for String conversions")
		}
	})
}

func TestBool(t *testing.T) {
	var c Conv
	t.Run("Bool", func(t *testing.T) {
		if n := assertions.EachOf(reflect.Bool, func(a *Assertion, e Expecter) {
			if err := e.Expect(c.Bool(a.From)); err != nil {
				t.Fatalf("%v:\n  %v", a.String(), err)
			}
		}); n < 1 {
			t.Fatalf("no test coverage ran for Bool conversions")
		}
	})
	t.Run("convNumToBool", func(t *testing.T) {
		var val reflect.Value
		if got, ok := c.convNumToBool(0, val); ok || got {
			t.Fatal("expected failure")
		}
	})
}

func TestBounds(t *testing.T) {
	defer initIntSizes(mathIntSize)

	var c Conv
	chk := func() {
		chkMaxInt, err := c.Int(fmt.Sprintf("%v", math.MaxInt64))
		if err != nil {
			t.Error(err)
		}
		if int64(chkMaxInt) != mathMaxInt {
			t.Errorf("chkMaxInt exp %v; got %v", chkMaxInt, mathMaxInt)
		}

		chkMinInt, err := c.Int(fmt.Sprintf("%v", math.MinInt64))
		if err != nil {
			t.Error(err)
		}
		if int64(chkMinInt) != mathMinInt {
			t.Errorf("chkMaxInt exp %v; got %v", chkMinInt, mathMaxInt)
		}

		chkUint, err := c.Uint(fmt.Sprintf("%v", uint64(math.MaxUint64)))
		if err != nil {
			t.Error(err)
		}
		if uint64(chkUint) != mathMaxUint {
			t.Errorf("chkMaxInt exp %v; got %v", chkMinInt, chkUint)
		}
	}

	initIntSizes(32)
	chk()

	initIntSizes(64)
	chk()
}

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

type ReadmeExample struct {
	Example *doc.Example
	Title   string
	Summary string
	Code    string
	Output  string
}

type ReadmeExamples []ReadmeExample

func (r ReadmeExamples) Len() int           { return len(r) }
func (r ReadmeExamples) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ReadmeExamples) Less(i, j int) bool { return r[i].Example.Order < r[j].Example.Order }

func TestReadmeGen(t *testing.T) {
	rewrap := func(buf *bytes.Buffer, with, s string) string {
		buf.Reset()
		for i := 1; i < len(s); i++ {
			if s[i-1] == 0xA {
				if s[i] == 0x9 {
					i++
				}
				buf.WriteString(with)
				continue
			}
			buf.WriteByte(s[i-1])
		}
		return buf.String()
	}

	sg := NewSrcGen("README.md")
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, "example_test.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	examples := doc.Examples(astFile)
	readmeExamples := make(ReadmeExamples, len(examples))

	for i, example := range examples {
		buf.Reset()
		format.Node(&buf, fset, example.Play)

		play := buf.String()
		idx := strings.Index(play, "func main() {") + 16
		if idx == -1 {
			t.Fatalf("bad formatting in example %v, could not find main() func", example.Name)
		}

		title := strings.Title(strings.TrimLeft(example.Name, "_"))
		if 0 == len(title) {
			title = "Package"
		}

		code := rewrap(&buf, "\n  > ", play[idx:len(play)-4])
		if 0 == len(code) {
			t.Fatalf("bad formatting in example %v, had no code", example.Name)
		}

		output := rewrap(&buf, "\n  > ", example.Output)
		if 0 == len(output) {
			t.Fatalf("bad formatting in example %v, had no output", example.Name)
		}

		summary := rewrap(&buf, "\n  ", example.Doc)
		if 0 == len(summary) {
			t.Fatalf("bad formatting in example %v, had no summary", example.Name)
		}

		readmeExamples[i] = ReadmeExample{
			Example: example,
			Title:   title,
			Summary: summary,
			Code:    code,
			Output:  output,
		}
	}

	sort.Sort(readmeExamples)
	sg.FuncMap["Examples"] = func() []ReadmeExample {
		return readmeExamples
	}

	if err := sg.Run(); err != nil {
		t.Fatal(err)
	}
}

type SrcGen struct {
	t           *testing.T
	Disabled    bool
	Data        interface{}
	FuncMap     template.FuncMap
	Name        string
	SrcPath     string
	SrcBytes    []byte
	TplPath     string
	TplBytes    []byte // Actual template bytes
	TplGenBytes []byte // Bytes produced from executed template
}

func NewSrcGen(name string) *SrcGen {
	funcMap := make(template.FuncMap)
	funcMap["args"] = func(s ...interface{}) interface{} {
		return s
	}
	funcMap["TrimSpace"] = strings.TrimSpace
	return &SrcGen{Name: "README.md", FuncMap: funcMap}
}

func (g *SrcGen) Run() error {
	if g.Disabled {
		return fmt.Errorf(`error: run failed because Disabled field is set for "%s"`, g.Name)
	}
	firstErr := func(funcs ...func() error) (err error) {
		for _, f := range funcs {
			err = f()
			if err != nil {
				return
			}
		}
		return
	}
	return firstErr(g.Check, g.Load, g.Generate, g.Format, g.Commit)
}

func (g *SrcGen) Check() error {
	g.Name = strings.TrimSpace(g.Name)
	g.TplPath = strings.TrimSpace(g.TplPath)
	g.SrcPath = strings.TrimSpace(g.SrcPath)
	if len(g.Name) == 0 {
		return errors.New("error: check for Name field failed because it was empty")
	}
	if len(g.TplPath) == 0 {
		g.TplPath = fmt.Sprintf(`testdata/%s.tpl`, g.Name)
	}
	if len(g.SrcPath) == 0 {
		g.SrcPath = fmt.Sprintf(`%s`, g.Name)
	}
	return nil
}

func (g *SrcGen) Load() error {
	var err error
	if g.TplBytes, err = ioutil.ReadFile(g.TplPath); err != nil {
		return fmt.Errorf(`error: load io error "%s" reading TplPath "%s"`, err, g.TplPath)
	}
	if g.SrcBytes, err = ioutil.ReadFile(g.SrcPath); err != nil {
		return fmt.Errorf(`error: load io error "%s" reading SrcPath "%s"`, err, g.SrcPath)
	}
	return nil
}

func (g *SrcGen) Generate() error {
	tpl, err := template.New(g.Name).Funcs(g.FuncMap).Parse(string(g.TplBytes))
	if err != nil {
		return fmt.Errorf(`error: generate error "%s" parsing TplPath "%s"`, err, g.TplPath)
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, g.Data); err != nil {
		return fmt.Errorf(`error: generate error "%s" executing TplPath "%s"`, err, g.TplPath)
	}
	g.TplGenBytes = buf.Bytes()
	return nil
}

func (g *SrcGen) Format() error {

	// Only run gofmt for .go source code.
	if !strings.HasSuffix(g.SrcPath, ".go") {
		return nil
	}

	fmtBytes, err := format.Source(g.TplGenBytes)
	if err != nil {
		return fmt.Errorf(`error: format error "%s" executing TplPath "%s"`, err, g.TplPath)
	}
	g.TplGenBytes = fmtBytes
	return err
}

func (g *SrcGen) IsStale() bool {
	return !bytes.Equal(g.SrcBytes, g.TplGenBytes)
}

func (g *SrcGen) Dump(w io.Writer) string {
	sep := strings.Repeat("-", 80)
	fmt.Fprintf(w, "%[1]s\n  TplBytes:\n%[1]s\n%s\n%[1]s\n", sep, g.TplBytes)
	fmt.Fprintf(w, "  SrcBytes:\n%[1]s\n%s\n%[1]s\n", sep, g.SrcBytes)
	fmt.Fprintf(w, "  TplGenBytes (IsStale: %v):\n%s\n%[3]s\n%[2]s\n",
		g.IsStale(), sep, g.TplGenBytes)
	return g.Name
}

func (g *SrcGen) String() string {
	return g.Name
}

func (g *SrcGen) Commit() error {
	if !g.IsStale() {
		return nil
	}
	return ioutil.WriteFile(g.SrcPath, g.TplGenBytes, 0644)
}

var refKindNames = map[reflect.Kind]string{
	DurationKind: "time.Duration",
	TimeKind:     "time.Time",
}

func reflectHasMask(k, mask reflect.Kind) bool {
	return (k & mask) == mask
}

func reflectKindStr(k reflect.Kind) string {
	if reflectHasMask(k, ChanMask|SliceMask|MapMask) {
		return refKindNames[k]
	}
	return k.String()
}

// returns the reflect.Kind with the additional masks / types for conversions.
func convKindStr(v interface{}) string { return reflectKindStr(convKind(v)) }
func convKind(v interface{}) reflect.Kind {
	switch T := indirect(v).(type) {
	case reflect.Kind:
		return T
	case Expecter:
		return T.Kind()
	case time.Time:
		return TimeKind
	case time.Duration:
		return DurationKind
	default:
		kind := reflect.TypeOf(T).Kind()
		if reflect.Slice == kind || kind == reflect.Array {
			kind |= SliceMask
		} else if kind == reflect.Map {
			kind |= MapMask
		} else if kind == reflect.Chan {
			kind |= SliceMask
		}
		return kind
	}
}

// Expecter defines a type of conversion for v and will return an error if the
// value is unexpected.
type Expecter interface {

	// Exp is given the result of conversion and an error if one occurred. The
	// type of conversion operation should be the type of Kind().
	Expect(got interface{}, err error) error

	// Kind returns the type of conversion that should be given to Exepct().
	Kind() reflect.Kind
}

type AssertionsLookup struct {
	assertion *Assertion
	expecter  Expecter
}

// // Assertions is map of assertions by the Expecter kind.
type AssertionsIndex map[reflect.Kind][]AssertionsLookup

// Assertions is slice of assertions.
type Assertions []*Assertion

// EachOf will visit each assertion that contains an Expecter for reflect.Kind.
func (a Assertions) EachOf(k reflect.Kind, f func(a *Assertion, e Expecter)) int {
	found := assertionsIndex[k]
	for _, v := range found {
		f(v.assertion, v.expecter)
	}
	return len(found)
}

// Assertion represents a set of expected values from a Converter. It's the
// general implementation of asserter. It will compare each Expects to the
// associated Converter.T(). It returns an error for each failed conversion
// or an empty error slice if no failures occurred.
type Assertion struct {

	// From is the underlying value to be used in conversions.
	From interface{}

	// Name of this assertion
	Name string

	// Code to construct Interface
	Code string

	// Type of Interface field using fmt %#T.
	Type string

	// File is the file the assertion is defined in.
	File string

	// Line is the line number the assertion is defined in.
	Line int

	// Description for this assertion.
	Desc string

	// Expects contains a list of values this Assertion expects to see as a result
	// of calling the conversion function for the associated type. For example
	// appending a single int64 to this slice would mean it is expected that:
	//   conv.Int64(assertion.From) == Expect[0].(int64)
	// When appending multiple values of the same type, the last one appended is
	// used.
	Exps []Expecter
}

func (a Assertion) String() string {
	return fmt.Sprintf("[assertion %v:%d] from `%v` (%[3]T)",
		a.File, a.Line, a.From)
}

// assert will create an assertion type for the first argument given. Args is
// consumed into the Expects slice, with slices of interfaces being flattened.
func assert(value interface{}, args ...interface{}) {
	a := &Assertion{
		From: value,
		Type: typename(value),
	}
	split := strings.SplitN(a.Type, ".", 2)
	a.Name = strings.Title(split[len(split)-1])
	a.Code = fmt.Sprintf("%#v", value)
	_, a.File, a.Line, _ = runtime.Caller(1)
	a.File = path.Base(a.File)

	var slurp func(interface{})
	slurp = func(value interface{}) {
		switch T := value.(type) {
		case []interface{}:
			for _, t := range T {
				slurp(t)
			}
		case []Expecter:
			for _, t := range T {
				slurp(t)
			}
		case Expecter:
			a.Exps = append(a.Exps, T)
		default:
			a.Exps = append(a.Exps, Exp{value})
		}
	}
	for _, arg := range args {
		slurp(arg)
	}
	for _, exp := range a.Exps {
		k := exp.Kind()
		assertionsIndex[k] = append(assertionsIndex[k], AssertionsLookup{a, exp})
	}
	assertions = append(assertions, a)
}

// Exp is used for ducmentation purposes.
type Exp struct {
	Want interface{}
}

func (e Exp) Kind() reflect.Kind {
	return convKind(e.Want)
}

func (e Exp) Expect(got interface{}, err error) error {
	if err == nil && !reflect.DeepEqual(got, e.Want) {
		return fmt.Errorf("(%T) %[1]v != %v (%[2]T)", got, e.Want)
	}
	return err
}

func (e Exp) String() string {
	return fmt.Sprintf("Exp{Want: %v (type %[1]T)}", e.Want)
}

// SkipExp is used for ducmentation purposes.
type SkipExp string

func (e SkipExp) Kind() reflect.Kind { return reflect.Invalid }
func (e SkipExp) Expect(interface{}, error) error {
	return nil
}

// FuncExp is expected to have it's own self contained test.
type FuncExp func(interface{}, error) error

func (e FuncExp) Kind() reflect.Kind { return reflect.Invalid }
func (e FuncExp) Expect(got interface{}, err error) error {
	return e(got, err)
}

// ErrorExp ensures a given conversion failed.
type ErrorExp struct {
	Exp
	ErrStr string
}

func experr(want interface{}, contains string) Expecter {
	return ErrorExp{ErrStr: contains, Exp: Exp{Want: want}}
}

func (e ErrorExp) Expect(got interface{}, err error) error {
	if err != nil {
		if len(e.ErrStr) == 0 {
			return err
		}
		if !strings.Contains(err.Error(), e.ErrStr) {
			return fmt.Errorf("error did not match:\n  exp: %v\n  got: %v", e.ErrStr, err)
		}
	} else if len(e.ErrStr) > 0 {
		return errors.New("expected non-nil err")
	}
	return nil
}

// Float64Exp asserts that (converter.Float64() - Exp) < epsilon64
type Float64Exp struct {
	Want float64
}

const (
	epsilon64 = float64(.00000000000000001)
)

func (e Float64Exp) Kind() reflect.Kind { return reflect.Float64 }
func (e Float64Exp) Expect(got interface{}, err error) error {
	val := got.(float64)
	abs := math.Abs(val - e.Want)
	if abs < epsilon64 {
		return err
	}
	return fmt.Errorf("%#v.assert(%#v): abs value %v exceeded epsilon %v",
		e, val, abs, epsilon64)
}

// Float32Exp asserts that (converter.Float32() - Exp) < epsilon64
type Float32Exp struct {
	Want float32
}

func (e Float32Exp) Kind() reflect.Kind { return reflect.Float32 }
func (e Float32Exp) Expect(got interface{}, err error) error {
	val := got.(float32)
	abs := math.Abs(float64(val - e.Want))
	if abs < epsilon64 {
		return err
	}
	return fmt.Errorf("%#v.assert(%#v): abs value %v exceeded epsilon %v",
		e, val, abs, epsilon64)
}

// TimeExp helps validate time.Time() conversions, specifically because under
// some conversions time.Now() may be used. It will check that the difference
// between the Moment is the same as the given value after truncation and
// rounding (if either is set) is identical.
type TimeExp struct {
	Moment   time.Time
	Offset   time.Duration
	Round    time.Duration
	Truncate time.Duration
}

func (e TimeExp) String() string     { return fmt.Sprintf("%v", e.Moment) }
func (e TimeExp) Kind() reflect.Kind { return TimeKind }
func (e TimeExp) Expect(got interface{}, err error) error {
	val := got.(time.Time).Add(e.Offset)
	if e.Round != 0 {
		val = val.Round(e.Round)
	}
	if e.Truncate != 0 {
		val = val.Round(e.Truncate)
	}
	if !e.Moment.Equal(val) {
		return fmt.Errorf(
			"times did not match:\n  exp: %v\n  got: %v", e.Moment, val)
	}
	return nil
}

// DurationExp supports fuzzy duration conversions.
type DurationExp struct {
	Want  time.Duration
	Round time.Duration
}

func (e DurationExp) Kind() reflect.Kind { return DurationKind }
func (e DurationExp) Expect(got interface{}, err error) error {
	d := got.(time.Duration)
	if e.Round != 0 {
		neg := d < 0
		if d < 0 {
			d = -d
		}
		if m := d % e.Round; m+m < e.Round {
			d = d - m
		} else {
			d = d + e.Round - m
		}
		if neg {
			d = -d
		}
	}
	if e.Want != d {
		return fmt.Errorf("%#v: %#v != %#v", e, e.Want, d)
	}
	return nil
}

// NowExp is like TimeExp but makes `Moment` time.Now
type NowExp struct {
	TimeExp
}

func (e NowExp) Kind() reflect.Kind { return TimeKind }
func (e NowExp) Expect(got interface{}, err error) error {
	e.TimeExp.Moment = time.Now()
	return e.TimeExp.Expect(got, err)
}

func typename(value interface{}) (name string) {
	parts := strings.SplitN(fmt.Sprintf("%T", value), ".", 2)

	if len(parts) == 2 {
		name = parts[0] + "." + parts[1]
	} else {
		name = parts[0]
	}
	return
}

type errHookConv struct {
	c  Converter
	fn func(from interface{}, err error) error
}

func (c errHookConv) Map(into, from interface{}) error {
	return c.fn(from, c.c.Map(into, from))
}

func (c errHookConv) Slice(into, from interface{}) error {
	return c.fn(from, c.c.Slice(into, from))
}

func (c errHookConv) Bool(from interface{}) (bool, error) {
	res, err := c.c.Bool(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Duration(from interface{}) (time.Duration, error) {
	res, err := c.c.Duration(from)
	return res, c.fn(from, err)
}

func (c errHookConv) String(from interface{}) (string, error) {
	res, err := c.c.String(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Time(from interface{}) (time.Time, error) {
	res, err := c.c.Time(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Float32(from interface{}) (float32, error) {
	res, err := c.c.Float32(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Float64(from interface{}) (float64, error) {
	res, err := c.c.Float64(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Int(from interface{}) (int, error) {
	res, err := c.c.Int(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Int8(from interface{}) (int8, error) {
	res, err := c.c.Int8(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Int16(from interface{}) (int16, error) {
	res, err := c.c.Int16(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Int32(from interface{}) (int32, error) {
	res, err := c.c.Int32(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Int64(from interface{}) (int64, error) {
	res, err := c.c.Int64(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Uint(from interface{}) (uint, error) {
	res, err := c.c.Uint(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Uint8(from interface{}) (uint8, error) {
	res, err := c.c.Uint8(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Uint16(from interface{}) (uint16, error) {
	res, err := c.c.Uint16(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Uint32(from interface{}) (uint32, error) {
	res, err := c.c.Uint32(from)
	return res, c.fn(from, err)
}

func (c errHookConv) Uint64(from interface{}) (uint64, error) {
	res, err := c.c.Uint64(from)
	return res, c.fn(from, err)
}

type testBoolConverter bool

func (t testBoolConverter) Bool() (bool, error) {
	return !bool(t), nil
}

type testStringConverter string

func (t testStringConverter) String() (string, error) {
	return string(t) + "Tested", nil
}

func init() {

	// bools
	{

		// strings: truthy
		trueStrings := []string{
			"1", "t", "T", "true", "True", "TRUE", "y", "Y", "yes", "Yes", "YES"}
		for _, truthy := range trueStrings {
			assert(truthy, true)
			assert(testStringConverter(truthy), true)
		}

		// strings: falsy
		falseStrings := []string{
			"0", "f", "F", "false", "False", "FALSE", "n", "N", "no", "No", "NO"}
		for _, falsy := range falseStrings {
			assert(falsy, false)
			assert(testStringConverter(falsy), false)
		}

		// numerics: true
		for _, i := range []int{-1, 1} {
			assert(int(i), true)
			assert(int8(i), true)
			assert(int16(i), true)
			assert(int32(i), true)
			assert(int64(i), true)
			assert(uint(i), true)
			assert(uint8(i), true)
			assert(uint16(i), true)
			assert(uint32(i), true)
			assert(uint64(i), true)
			assert(float32(i), true)
			assert(float64(i), true)
			assert(complex(float32(i), 0), true)
			assert(complex(float64(i), 0), true)
		}

		// int/uint: false
		assert(int(0), false)
		assert(int8(0), false)
		assert(int16(0), false)
		assert(int32(0), false)
		assert(int64(0), false)
		assert(uint(0), false)
		assert(uint8(0), false)
		assert(uint16(0), false)
		assert(uint32(0), false)
		assert(uint64(0), false)

		// float: NaN and 0 are false.
		assert(float32(math.NaN()), false)
		assert(math.NaN(), false)
		assert(float32(0), false)
		assert(float64(0), false)

		// complex: NaN and 0 are false.
		assert(complex64(cmplx.NaN()), false)
		assert(cmplx.NaN(), false)
		assert(complex(float32(0), 0), false)
		assert(complex(float64(0), 0), false)

		// time
		assert(time.Time{}, false)
		assert(time.Now(), true)

		// bool
		assert(false, false)
		assert(true, true)

		// underlying bool
		type ulyBool bool
		assert(ulyBool(false), false)
		assert(ulyBool(true), true)

		// implements bool converter
		assert(testBoolConverter(false), true)
		assert(testBoolConverter(true), false)

		// test length kinds
		assert([]string{"one", "two"}, true)
		assert(map[int]string{1: "one", 2: "two"}, true)
		assert([]string{}, false)
		assert([]string(nil), false)

		// errors
		assert(nil, experr(false, `cannot convert <nil> (type <nil>) to bool`))
		assert("foo", experr(false, `cannot parse "foo" (type string) as bool`))
		assert("tooLong", experr(
			false, `cannot parse type string with len 7 as bool`))
		assert(struct{}{}, experr(
			false, `cannot convert struct {}{} (type struct {}) to `))
	}

	// strings
	{

		// basic
		assert(`hello`, `hello`)
		assert(``, ``)
		assert([]byte(`hello`), `hello`)
		assert([]byte(``), ``)

		// ptr indirection
		assert(new(string), ``)
		assert(new([]byte), ``)

		// underlying string
		type ulyString string
		assert(ulyString(`hello`), `hello`)
		assert(ulyString(``), ``)

		// implements string converter
		assert(testStringConverter(`hello`), `helloTested`)
		assert(testStringConverter(`hello`), `helloTested`)

		// errors
		assert(nil, experr(false, `cannot convert <nil> (type <nil>) to bool`))
		assert("foo", experr(false, `cannot parse "foo" (type string) as bool`))
		assert("tooLong", experr(
			false, `cannot parse type string with len 7 as bool`))
		assert(struct{}{}, experr(
			false, `cannot convert struct {}{} (type struct {}) to `))
	}
}
