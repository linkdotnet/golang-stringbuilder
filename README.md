[![Go](https://github.com/linkdotnet/golang-stringbuilder/actions/workflows/go.yml/badge.svg)](https://github.com/linkdotnet/golang-stringbuilder/actions/workflows/go.yml)
[![Go Report card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)
[![GoDoc](https://pkg.go.dev/badge/github.com/linkdotnet/golang-stringbuilder?status.svg)](https://pkg.go.dev/github.com/linkdotnet/golang-stringbuilder?tab=doc)

# golang-stringbuilder
A string builder that has similar capabilities as the one from C#. The goal is to have a straightforward API that lets you work with strings easily.

## Install

To install the package, call:
```bash
go get -u github.com/linkdotnet/golang-stringbuilder
```

Next import the package:
```golang
import ( "github.com/linkdotnet/golang-stringbuilder" )
```

## Quickstart

The API derives from the C# `StringBuilder`. You can easily append strings or single runes.

```golang
func main() {
	sb := Text.StringBuilder{}
	sb.Append("Hello")
	sb.Append(" ")
	sb.Append("World")
	fmt.Println(sb.ToString())
}
```

Also more advanced use cases where you want to insert an arbitrary word at an arbitrary position are possible.
```golang
sb := NewStringBuilderFromString("Hello World")
sb.Insert(5, " my dear")
output := sb.ToString() // Hello my dear World
```

The `StringBuilder` also implements the `io.Writer` interface so it can be use with `fmt.Fprintf` and friends:
```golang
s := &StringBuilder{}

for i := 3; i >= 1; i-- {
	fmt.Fprintf(s, "%d...", i)
}

s.Append("lift off")
fmt.PrintLn(s.ToString()) // Prints 3...2...1...lift off
```

## Benchmark
Check out the implementation of the benchmark in the corresponding file. Here are some results:
```no-class
goos: darwin
goarch: arm64
pkg: github.com/linkdotnet/golang-stringbuilder
BenchmarkStringBuilderConcat
BenchmarkStringBuilderConcat-10           343710              3394 ns/op            4352 B/op          5 allocs/op
BenchmarkStringConcat
BenchmarkStringConcat-10                 1000000              1113 ns/op            6720 B/op         24 allocs/op
BenchmarkGoStringBuilderConcat
BenchmarkGoStringBuilderConcat-10        4204142               278.4 ns/op          1448 B/op          6 allocs/op
```