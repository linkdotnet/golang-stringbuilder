package Text

import "testing"

func TestAppend(t *testing.T) {
	const expected string = "Hello World"
	sb := StringBuilder{}

	sb.Append(expected)

	if result := sb.ToString(); result != expected {
		t.Errorf("Actual %q, Expected: %q", result, expected)
	}
}

func TestLen(t *testing.T) {
	sb := StringBuilder{}

	sb.Append("1234")

	if len := sb.Len(); len != 4 {
		t.Errorf("Actual %q, Expected: %q", len, 4)
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
