# Go Package: conv

  [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/cstockton/go-conv)
  [![Go Report Card](https://goreportcard.com/badge/github.com/cstockton/go-conv?style=flat-square)](https://goreportcard.com/report/github.com/cstockton/go-conv)
  [![Coverage Status](https://img.shields.io/codecov/c/github/cstockton/go-conv/master.svg?style=flat-square)](https://codecov.io/github/cstockton/go-conv?branch=master)
  [![Build Status](http://img.shields.io/travis/cstockton/go-conv.svg?style=flat-square)](https://travis-ci.org/cstockton/go-conv)
  [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/cstockton/go-conv/master/LICENSE)

  > Get:
  > ```bash
  > go get -u github.com/cstockton/go-conv
  > ```
  >
  > Example:
  > ```Go
  > // Basic types
  > b := conv.Bool("YES")
  >   b -> true
  >
  > // Slices
  > var into []int64
  > err := conv.Slice(&into, []string{"12", "345", "6789"})
  >   into -> []int64{12, 234, 6789}
  >
  > // Maps
  > into := make(map[string]int64)
  > err := conv.Map(into, []string{"12", "345", "6789"})
  >   into -> map[string]int64{"0": 12, "1", 234, "2", 6789}
  > ```


## Intro

Package conv provides fast and intuitive conversions across Go types. This library uses reflection to be robust but will bypass it for common conversions, for example string conversion to any type will never use reflection. In most cases this library is as fast or faster then the standard library for similar operations due to various aggressive (but safe) optimizations. The only external dependency ([iter](https://github.com/cstockton/go-iter)) has 100% test coverage and is maintained by me. It is used to walk the values given for map and slice conversion and will **never panic**. All methods and functions are **safe for concurrent use by multiple Goroutines**, with a single exception that Slice and Map conversion under certain circumstances may produce undefined results if they are mutated while being traversed.


{{ range Examples -}}
### {{ .Title }}

  {{ .Summary }}

  > Example:
  > ```Go
  > {{ .Code }}
  > ```
  >
  > Output:
  > ```Go
  > {{ .Output }}
  > ```


{{ end -}}

## Contributing

Feel free to create issues for bugs, please ensure code coverage remains 100%
with any pull requests.


## Bugs and Patches

  Feel free to report bugs and submit pull requests.

  * bugs:
    <https://github.com/cstockton/go-conv/issues>
  * patches:
    <https://github.com/cstockton/go-conv/pulls>
