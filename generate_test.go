package conv

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"time"
)

func TestGenerate(t *testing.T) {
	for _, generator := range generators {
		generator := generator

		t.Run(generator.name, func(t *testing.T) {
			generator.Run(t)
		})
	}
}

func TestUtils(t *testing.T) {
	t.Run("CodeFuncs", func(t *testing.T) {
		chk := func(using func(given interface{}) string, given interface{}, want string) {
			if got := using(given); got != want {
				t.Fatalf("given: %v got: %v want: %v", given, got, want)
			}
		}
		chk(codeValue, "123", `"123"`)
		chk(codeValue, time.Duration(time.Second*10), `time.Duration(10000000000)`)
		chk(codeValue, time.Time{}, `time.Time{sec:0, nsec:0, loc: time.UTC}`)
		chk(codeValue, time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC),
			`time.Time{sec:63393490862, nsec:3, loc: time.UTC}`)
		chk(codeType, DocAssertion("123"), `123`)
		chk(codeType, "123", `string`)
		chk(codeType, time.Duration(time.Second*10), `time.Duration`)
		chk(codeType, time.Time{}, `time.Time`)
		chk(codeConvFuncName, DocAssertion("123"), `123`)
		chk(codeConvFuncName, "123", `String`)
		chk(codeConvFuncName, time.Duration(time.Second*10), `Duration`)
		chk(codeConvFuncName, time.Time{}, `Time`)
	})
	t.Run("CodeConvFunc", func(t *testing.T) {
		chk := func(from, to interface{}, want string) {
			if got := codeConvFunc(from, to); got != want {
				t.Fatalf("from: %v to: %v\n  got: %v\n  want: %v", from, to, got, want)
			}
		}
		chk(1234, int(1234), `Int(1234)`)
		chk("123ns", time.Duration(123), `Duration("123ns")`)
		chk("foo", time.Time{}, `Time("foo")`)
		chk(DocAssertion(`Docs(123)`), time.Time{}, `Docs(123)`)
	})
}

var generators = []*Generator{
	NewGenerator(`README.md`),
	NewGenerator(`doc.go`),
	NewGenerator(`numeric.go`),
}

func codeValue(v interface{}) string {
	datefmt := func(v interface{}) string {
		val := fmt.Sprintf("%#v", v)
		val = strings.Split(val, ", loc")[0]
		return fmt.Sprintf("%v, loc: time.UTC}", val)
	}
	switch T := Indirect(v).(type) {
	case DocAssertion:
		return string(T)
	case TimeAssertion:
		return datefmt(T.Moment)
	case NowAssertion:
		return datefmt(time.Now().Add(T.Offset))
	case time.Time:
		return datefmt(T)
	case time.Duration:
		return fmt.Sprintf("time.Duration(%#v)", T)
	}
	return fmt.Sprintf("%#v", v)
}

func codeType(v interface{}) string {
	switch T := Indirect(v).(type) {
	case DocAssertion:
		return string(T)
	}
	return fmt.Sprintf("%T", v)
}

func codeConvFuncName(v interface{}) string {
	switch T := Indirect(v).(type) {
	case DocAssertion:
		return string(T)
	case time.Time, TimeAssertion, NowAssertion:
		return `Time`
	case time.Duration:
		return `Duration`
	}
	return strings.Title(fmt.Sprintf("%T", v))
}

func codeConvFunc(from, to interface{}) string {
	switch T := Indirect(from).(type) {
	case DocAssertion:
		return string(T)
	}
	return fmt.Sprintf("%v(%v)", codeConvFuncName(to), codeValue(from))
}

type Generator struct {
	name         string
	data         interface{}
	funcs        template.FuncMap
	source       []byte
	target       []byte
	generated    []byte
	generatedSum [md5.Size]byte
	targetSum    [md5.Size]byte
	funcsMap     template.FuncMap
}

func NewGenerator(name string) *Generator {
	funcMap := make(template.FuncMap)
	funcMap["ToLower"] = strings.ToLower
	funcMap["ToUpper"] = strings.ToUpper
	funcMap["Title"] = func(s interface{}) string {
		return strings.Title(fmt.Sprintf("%v", s))
	}
	funcMap["Convert"] = func(from, to, name string) string {
		if from == to {
			return name
		}
		return fmt.Sprintf("%v(%v)", to, name)
	}
	funcMap["args"] = func(s ...interface{}) interface{} {
		return s
	}
	funcMap["Summary"] = func() string {
		return summary
	}
	funcMap["Sections"] = func() []Section {
		return sections
	}
	funcMap["Assertions"] = func() Assertions {
		return assertions
	}
	funcMap["CodeValue"] = codeValue
	funcMap["CodeType"] = codeType
	funcMap["CodeConvFuncName"] = codeConvFuncName
	funcMap["CodeConvFunc"] = codeConvFunc
	funcMap["Bool"] = Bool
	funcMap["Complex64"] = Complex64
	funcMap["Complex128"] = Complex128
	funcMap["Duration"] = Duration
	funcMap["Float32"] = Float32
	funcMap["Float64"] = Float64
	funcMap["Int"] = Int
	funcMap["Int8"] = Int8
	funcMap["Int16"] = Int16
	funcMap["Int32"] = Int32
	funcMap["Int64"] = Int64
	funcMap["String"] = String
	funcMap["Time"] = Time
	funcMap["Uint"] = Uint
	funcMap["Uint8"] = Uint8
	funcMap["Uint16"] = Uint16
	funcMap["Uint32"] = Uint32
	funcMap["Uint64"] = Uint64
	return &Generator{name: name, funcs: funcMap}
}

func (g *Generator) Run(t *testing.T) {
	g.check(t)
	g.load(t)
	g.generate(t)
	g.commit(t)
}

func (g *Generator) check(t *testing.T) {
	abs, err := filepath.Abs(g.name)
	if err != nil {
		t.Errorf("error[%s] checking abs path for target file: %s", err, g.name)
	}
	if !strings.HasSuffix(abs, filepath.Join(`go-conv`, g.name)) {
		t.Errorf("error[%s] checking suffix target file: %s", err, g.name)
	}
}

func (g *Generator) load(t *testing.T) {
	p := filepath.Join(`testdata`, g.name+`.tpl`)
	var err error
	if g.source, err = ioutil.ReadFile(p); err != nil {
		t.Errorf("error[%s] reading target file: %s", err, p)
	}
	if g.target, err = ioutil.ReadFile(g.name); err != nil {
		t.Errorf("error[%s] reading target file: %s", err, g.name)
	}
	g.targetSum = md5.Sum(g.target)
}

func (g *Generator) generate(t *testing.T) {
	tpl, err := template.New(g.name).
		Funcs(g.funcs).
		Parse(string(g.source))
	if err != nil {
		t.Fatalf("error[%s] parsing source template: %s", err, g.name)
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, nil); err != nil {
		t.Fatalf("error[%s] executing source template: %s", err, g.name)
	}
	g.generated = g.format(t, buf.Bytes())
	g.generatedSum = md5.Sum(g.generated)
}

func (g *Generator) format(t *testing.T, b []byte) []byte {
	if !strings.HasSuffix(g.name, ".go") || g.name == "doc.go" {
		t.Logf("info[%s] skipping go fmt", g.name)
		return b
	}
	var (
		sout bytes.Buffer
		serr bytes.Buffer
	)
	cmd := exec.Command("gofmt")
	cmd.Stdin = bytes.NewReader(b)
	cmd.Stdout = &sout
	cmd.Stderr = &serr

	err := cmd.Run()
	if len(serr.String()) > 0 {
		for _, val := range strings.Split(serr.String(), "\n") {
			if len(val) > 0 {
				t.Logf("error[%s] gofmt said: %v", g.name, val)
			}
		}
		t.Logf("error[%s] could not format file: %s",
			g.name, string(g.generated))
		t.Fatalf("error[%s] could not format file", g.name)
	}
	if err != nil {
		t.Fatalf("error[%s] could not run gofmt on file: %v", g.name, err)
	}
	return sout.Bytes()
}

func (g *Generator) commit(t *testing.T) {
	if g.generatedSum == g.targetSum {
		t.Skip("skipping because target file is identical to generated output.")
	}
	if !*generate {
		t.Fatalf("error[%s OUTDATED] run `go generate` to update it", g.name)
	}
	if *nowrite {
		fmt.Fprintf(os.Stdout, "%s\n", g.name)
		_, _ = io.Copy(os.Stdout, bytes.NewReader(g.generated))
		t.Skip("skipping writing because nowrite flag.")
	}
	if err := ioutil.WriteFile(g.name, g.generated, 0644); err != nil {
		t.Errorf("error[%s] writing to output file: %s", err, g.name)
	}
}
