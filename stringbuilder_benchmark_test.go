package Text

import (
	"strings"
	"testing"
)

const text = "Hello dear World"
const count = 25

var result string

func BenchmarkStringBuilderConcat(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = benchmarkStringBuilderConcat(text, count)
	}
	result = r
}

func BenchmarkStringConcat(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = benchmarkStringConcat(text, count)
	}
	result = r
}

func BenchmarkGoStringBuilderConcat(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = benchmarkGoStringBuilderConcat(text, count)
	}
	result = r
}

func benchmarkStringBuilderConcat(text string, count int) string {
	s := NewStringBuilder(64)
	for i := 0; i < count; i++ {
		s.Append(text)
	}

	return s.ToString()
}

func benchmarkStringConcat(text string, count int) string {
	s := ""
	for i := 0; i < count; i++ {
		s += text
	}

	return s
}

func benchmarkGoStringBuilderConcat(text string, count int) string {
	s := &strings.Builder{}

	for i := 0; i < count; i++ {
		s.WriteString(text)
	}

	return s.String()
}
