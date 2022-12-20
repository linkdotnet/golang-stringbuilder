# golang-stringbuilder
A string builder that has similar capabilities as the one from C#. The goal is to have a straightforward API that lets you work with strings easily.

## Install

To install the package, call:
```bash
go get -u github.com/linkdotnet/golang-stringbuilder 
```

## Usage
```golang
package sample

import (
	"fmt"

	Text "github.com/linkdotnet/golang-stringbuilder"
)

func main() {
	sb := Text.StringBuilder{}
	sb.Append("Hello")
	sb.Append(" ")
	sb.Append("World")
	fmt.Println(sb.ToString())
}

```