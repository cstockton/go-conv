# Go Package: conv

  <a href="https://godoc.org/github.com/cstockton/go-conv"><img src="https://img.shields.io/badge/%20docs-reference-5272B4.svg?style=flat-square"></a> [![Go Report Card](https://goreportcard.com/badge/github.com/cstockton/go-conv)](https://goreportcard.com/report/github.com/cstockton/go-conv)

  > Get:
  > ```bash
  > go get -u github.com/cstockton/go-conv
  > ```
  >
  > Example:
  > ```Go
  > // Basic types
  > b := conv.Bool("true")
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


## About

Small library to make working with various types of values a bit easier. See the
documentation for more information.


## Contributing

Feel free to create issues for bugs, please ensure code coverage remains 100%
with any pull requests.


## Bugs and Patches

  Feel free to report bugs and submit pull requests.

  * bugs:
    <https://github.com/cstockton/go-conv/issues>
  * patches:
    <https://github.com/cstockton/go-conv/pulls>
