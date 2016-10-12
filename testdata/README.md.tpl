# Go Package: conv

  [About](#about) | [Usage](#usage) | [Go Doc](https://godoc.org/github.com/cstockton/go-conv)

  > Get:
  > ```bash
  > go get -u github.com/cstockton/go-conv
  > ```
  >
  > Example:
  > ```Go
  > conv.Bool("true") // bool(true)
  > conv.Int64("-1.23") // -1
  > conv.Duration("123ns") // time.Duration(123)
  > ```


## About

Small library to make working with various types of values a bit easier. When
`go generate` was released I thought it would be useful to integrate into unit
tests and code generation for some of the more boiler plate type libraries
and operations I perform in Go, like conversions. All heavy unit testing types
as well as assertion definitions are within `*_test.go` files in the `init()`
function. This means they are not exported and you do not pay an allocation tax
at program startup for these features.

{{ template "usage" }}

{{- template "contributing" }}

{{- template "bugs" -}}


{{ define "contributing" }}
## Contributing

Feel free to make a issues defining some conversion's that don't exist or
challenging the current ones. If you find a conversion you wish to change the
behavior of you MUST add an Assertion within a `*_test` file. A good place may
be [conv_test.go](blob/master/conv_test.go).

  - Example assertions for a Complex128 type, the functions are helpers that are defined in [assert_test.go](blob/master/assert_test.go).

    > Go:
    > ```Go
    > asserts("Complex128",
    >   assert(complex128(-1), true, "(-1+0i)", numerics(-1), durations(-1)),
    >   assert(complex128(0), false, "(0+0i)", numerics(0), durations(0)),
    >   assert(complex128(1), true, "(1+0i)", numerics(1), durations(1)),
    > )

This may cause other documentation to need updating with the `go generate` command.
Keep in mind the documentation through the `Section` and `Group` types is meant
to be consumable rather than comprehensive. Most conversion Assertions should live
outside of those. The template files are in the [testdata](blob/master/testdata)
folder.

The source and target files are defined below:

  - [README.md](blob/master/README.md) -> [README.md.tpl](blob/master/testdata/README.md.tpl)
  - [doc.go](blob/master/doc.go) -> [README.md.tpl](blob/master/testdata/doc.go.tpl)
  - [numeric.go](blob/master/numeric.go) -> [README.md.tpl](blob/master/testdata/numeric.go.tpl)

{{ end -}}


{{ define "bugs" }}
## Bugs and Patches

  Feel free to report bugs and submit pull requests.

  * bugs:
    <https://github.com/cstockton/go-conv/issues>
  * patches:
    <https://github.com/cstockton/go-conv/pulls>



[Go Doc]: https://godoc.org/github.com/cstockton/go-value
{{ end -}}


{{ define "usage" }}
## Usage

{{ Summary }}

{{ range $section := Sections -}}
{{ template "section" $section -}}
{{ end }}
{{ end -}}

{{ define "section" -}}
### {{ .Name }}

{{ range $group := .Groups -}}
{{ .Description }}
{{ template "group" $group }}
{{- end }}
{{ end -}}


{{ define "group" }}
  > Example:
  > ```Go
{{- range $assertion := .Assertions }}
{{- if $assertion.Expects }}
{{- $expect := index ($assertion.Expects) 0 }}
  > {{ CodeConvFunc $assertion.Interface $expect }} // {{ CodeValue $expect }}
{{- else }}
  > {{ $assertion.Interface }}
{{- end }}
{{- end }}
  > ```

{{ end -}}
