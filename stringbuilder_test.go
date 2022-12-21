package Text

import (
	"testing"
)

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

func TestToStringEmptyBuilder(t *testing.T) {
	sb := StringBuilder{}

	if result := sb.ToString(); result != "" {
		t.Errorf("String should be empty but was: %q", result)
	}
}

func TestNewFromString(t *testing.T) {
	const expected string = "Hello"

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
