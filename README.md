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