package Text

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"Hello World"},
		{"Hallöchen"},
		{"汉字汉字汉字汉字汉字汉字"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			s := &StringBuilder{}

			s.Append(tt.want)

			if got := s.ToString(); got != tt.want {
				t.Errorf("StringBuilder.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendMultipleTypes(t *testing.T) {
	tests := []struct {
		stringInput     string
		intInput        int
		booleanInput    bool
		multipleStrings []string
		want            string
	}{
		{"hello", 123, false, []string{"a", "b", "c"}, "hello123falseabc"},
		{"hello", 123, true, []string{"a", "b", "c"}, "hello123trueabc"},
	}
	for _, tt := range tests {
		s := &StringBuilder{}
		s.Append(tt.stringInput).AppendInt(tt.intInput).AppendBool(tt.booleanInput).AppendList(tt.multipleStrings)

		if got := s.ToString(); got != tt.want {
			t.Errorf("StringBuilder.Append Multiple types = %v, want %v", got, tt.want)
		}
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"English word", "Hello", 5},
		{"Word with Umlaut", "Hallöchen", 9},
		{"Multibyte", "汉字", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StringBuilder{}

			s.Append(tt.input)

			if got := s.Len(); got != tt.want {
				t.Errorf("StringBuilder.Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendLine(t *testing.T) {
	const expected string = "Hello World\n"
	sb := NewStringBuilder(100)

	sb.AppendLine("Hello World")

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestAppendRune(t *testing.T) {
	const expected string = "Hello"
	sb := NewStringBuilder(100)

	sb.Append("Hell")
	sb.AppendRune('o')

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestToStringEmptyBuilder(t *testing.T) {
	sb := StringBuilder{}

	if result := sb.ToString(); result != "" {
		t.Errorf("String should be empty but was: %q", result)
	}
}

func TestNewFromString(t *testing.T) {
	const expected string = "Hellöchen"

	sb := NewStringBuilderFromString(expected)

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestRemovePartOfString(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	if err := sb.Remove(3, 2); err != nil {
		t.Errorf("Remove threw an error: %v", err)
	}

	if result := sb.ToString(); result != "Hel" {
		t.Errorf("Actual %q, Expected: %q", result, "Hel")
	}
}

func TestRemoveWhenStartIndexOutOfBounds(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	if err := sb.Remove(100, 1); err == nil {
		t.Error("Should throw error but did not")
	}
}

func TestRemoveWhenStartIndexNegative(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	if err := sb.Remove(-1, 1); err == nil {
		t.Error("Should throw error but did not")
	}
}

func TestRemoveWhenLengthNegative(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	if err := sb.Remove(1, -1); err == nil {
		t.Error("Should throw error but did not")
	}
}

func TestRemoveWhenEndIndexOutOfBounds(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	if err := sb.Remove(4, 4); err == nil {
		t.Error("Should throw error but did not")
	}
}

func TestRemoveWhenLengthZero(t *testing.T) {
	const expected string = "Hello"
	sb := NewStringBuilderFromString(expected)

	if err := sb.Remove(0, 0); err != nil {
		t.Errorf("Remove threw an error: %v", err)
	}

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestRemoveInTheMiddle(t *testing.T) {
	const expected string = "Hlo World"
	sb := NewStringBuilderFromString("Hello World")

	if err := sb.Remove(1, 2); err != nil {
		t.Errorf("Remove threw an error: %v", err)
	}

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestInsertAtIndex(t *testing.T) {
	const expected string = "Hello my dear and beautiful World"
	sb := NewStringBuilderFromString("Hello World")

	if err := sb.Insert(5, " my dear and beautiful"); err != nil {
		t.Errorf("Insert threw an error: %v", err)
	}

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestInsertShouldThrowIfNegativeIndex(t *testing.T) {
	sb := StringBuilder{}

	if err := sb.Insert(-1, "Test"); err == nil {
		t.Error("Should throw error but did not")
	}
}

func TestInsertShouldThrowErrirIfOutOfRange(t *testing.T) {
	sb := StringBuilder{}

	if err := sb.Insert(1, "Test"); err == nil {
		t.Error("Should throw error but did not")
	}
}

func TestClear(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	sb.Clear()

	if result := sb.ToString(); result != "" {
		t.Errorf("Expected empty string but did receive %v", result)
	}
}

func TestRuneAt(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	if result := sb.RuneAt(1); result != 'e' {
		t.Errorf("Actual %q, Expected: %q", result, 'e')
	}
}

func TestAsRune(t *testing.T) {
	expected := []rune{'H', 'e', 'l', 'l', 'o'}
	sb := NewStringBuilderFromString("Hello")

	if result := sb.AsRuneSlice(); !reflect.DeepEqual(result, expected) {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestAppendRuneMultiple(t *testing.T) {
	expected := "aaaaaaaaaaaaaaa"
	sb := StringBuilder{}

	for i := 0; i < 15; i++ {
		sb.AppendRune('a')
	}

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestFindFirst(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		needle string
		want   int
	}{
		{"Empty haystack", "", "n", -1},
		{"Empty needle", "n", "", -1},
		{"Needle longer than haystack", "a", "ab", -1},
		{"Hello in Hello World", "Hello World", "Hello", 0},
		{"ö in Helöö", "Hellöö", "ö", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStringBuilderFromString(tt.input)
			if got := s.FindFirst(tt.needle); got != tt.want {
				t.Errorf("StringBuilder.FindFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindLast(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		needle string
		want   int
	}{
		{"Empty haystack", "", "n", -1},
		{"Empty needle", "n", "", -1},
		{"Needle longer than haystack", "a", "ab", -1},
		{"Hello in Hello World", "Hello World", "Hello", 0},
		{"ö in Helöö", "Hellöö", "ö", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStringBuilderFromString(tt.input)
			if got := s.FindLast(tt.needle); got != tt.want {
				t.Errorf("StringBuilder.FindLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		needle string
		want   []int
	}{
		{"Empty haystack", "", "n", []int{}},
		{"Empty needle", "n", "", []int{}},
		{"Needle longer than haystack", "a", "ab", []int{}},
		{"Hello in Hello World", "Hello World", "Hello", []int{0}},
		{"ö in Helöö", "Hellöö", "ö", []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStringBuilderFromString(tt.input)
			if got := s.FindAll(tt.needle); !slicesEqual(got, tt.want) {
				t.Errorf("StringBuilder.FindLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceRune(t *testing.T) {
	s := NewStringBuilderFromString("Hello")

	s.ReplaceRune('l', 'm')

	if got := s.ToString(); got != "Hemmo" {
		t.Errorf("StringBuilder.ReplaceRune() = %v, want %v", got, "Hemmo")
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		oldValue string
		newvalue string
		want     string
	}{
		{"Replace Hello with Hallo", "Hello World", "Hello", "Hallo", "Hallo World"},
		{"Replace Hello with Ha", "Hello World", "Hello", "Ha", "Ha World"},
		{"Replace Hello with Hallochen", "Hello World", "Hello", "Hallochen", "Hallochen World"},
		{"Replace Hello with Hallöchen", "Hello World", "Hello", "Hallöchen", "Hallöchen World"},
		{"Replace ö with ä", "äö", "ö", "ä", "ää"},
		{"Replace with same word", "Hello", "llo", "llo", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStringBuilderFromString(tt.input)

			s.Replace(tt.oldValue, tt.newvalue)

			if got := s.ToString(); got != tt.want {
				t.Errorf("StringBuilder.FindLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	const want = "3...2...1..."
	s := &StringBuilder{}

	for i := 3; i >= 1; i-- {
		fmt.Fprintf(s, "%d...", i)
	}

	if got := s.ToString(); got != want {
		t.Errorf("StringBuilder.Write() = %v, want %v", got, want)
	}
}

func TestWriteReturnsAddedAmount(t *testing.T) {
	s := &StringBuilder{}

	if got, _ := fmt.Fprintf(s, "%v", "Hello"); got != 5 {
		t.Errorf("StringBuilder.Write() = %v, want %v", got, 5)
	}
}

func TestTrimStartWithWhitespaces(t *testing.T) {
	s := NewStringBuilderFromString("   Hello World")

	s.TrimStart()

	if got := s.ToString(); got != "Hello World" {
		t.Errorf("StringBuilder.TrimStart() = %v, want %v", got, "Hello World")
	}
}

func TestTrimStartWithGivenCharacters(t *testing.T) {
	s := NewStringBuilderFromString("aHello World")

	s.TrimStart('a')

	if got := s.ToString(); got != "Hello World" {
		t.Errorf("StringBuilder.TrimStart() = %v, want %v", got, "Hello World")
	}
}

func TestTrimWithWhitespacesAtTheStartAndEnd(t *testing.T) {
	s := NewStringBuilderFromString("   Hello World   ")

	s.Trim()

	if got := s.ToString(); got != "Hello World" {
		t.Errorf("StringBuilder.Trim() = %v, want %v", got, "Hello World")
	}
}

func TestReverseStringBuilder(t *testing.T) {
	tests := []struct {
		name   string
		insert []string
		want   string
	}{
		{"Reverse odd length string builder", []string{"A", "B", "C"}, "CBA"},
		{"Reverse even length string builder", []string{"A", "B"}, "BA"},
		{"Reverse string builder of size 1", []string{"A"}, "A"},
		{"Reverse empty string builder", []string{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := &StringBuilder{}

			for _, s := range tt.insert {
				sb = sb.Append(s)
			}

			sb = sb.Reverse()

			if got := sb.ToString(); got != tt.want {
				t.Errorf("StringBuilder.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReuseReversedStringBuilder(t *testing.T) {
	sb := StringBuilder{}

	sb = *sb.Append("A").Append("B").Append("C").Reverse().Append("X")
	if got := sb.ToString(); got != "CBAX" {
		t.Errorf("StringBuilder.Reverse() = %v, want %v", got, "CBAX")
	}
}

func TestStringBuilderSubstring(t *testing.T) {
	tests := []struct {
		name         string
		start        int
		end          int
		substring    string
		errorMessage string
	}{
		{"Substring with negative start", -1, 3, "", "start should always be greater than or equal to zero"},
		{"Substring with end exceeding string builder length", 0, 5, "", "end cannot be greater than the length of string builder"},
		{"Substring with start greater than end", 3, 2, "", "start cannot be greater than the end for Substring() function"},
		{"Substring with start equal to zero", 0, 3, "abc", ""},
		{"Substring with end equal to length of string builder", 0, 4, "abcd", ""},
		{"Substring of length 1", 0, 1, "a", ""},
		{"Substring of length 0", 0, 0, "", ""},
		{"Substring in middle of string builder", 1, 3, "bc", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := NewStringBuilderFromString("abcd")
			s, err := sb.Substring(tt.start, tt.end)
			if err != nil {
				if err.Error() != tt.errorMessage {
					t.Errorf("StringBuilder.Substring() expected error message = %v, got = %v", tt.errorMessage, err.Error())
				}

			}
			if s != tt.substring {
				t.Errorf("StringBuilder.Substring() expected substring = %v, got = %v", tt.substring, s)
			}
		})
	}
}

func slicesEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSetRuneAt(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	err := sb.SetRuneAt(10, 'a')
	if err == nil {
		t.Errorf("Should throw error but did not")
	}

	err = sb.SetRuneAt(0, 'a')
	if err != nil {
		t.Errorf("Should not throw error")
	}

	if result := sb.RuneAt(0); result != 'a' {
		t.Errorf("Actual %q, Expected: %q", result, 'e')
	}
}

func TestSkip(t *testing.T) {
	sb := NewStringBuilderFromString("Hello")

	err := sb.Skip(1)
	if err == nil {
		t.Errorf("Should throw error but did not")
	}

	sb.Clear()

	err = sb.Skip(1)
	if err != nil {
		t.Errorf("Should not throw error")
	}

	if sb.position != 1 {
		t.Errorf("Actual position %d, Expected: 1", sb.position)
	}
}
