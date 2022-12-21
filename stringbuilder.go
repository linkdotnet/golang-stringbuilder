package Text

import "fmt"

type StringBuilder struct {
	data     []rune
	position int
}

// Creates a new instance of the StringBuilder with preallocated array
func NewStringBuilder(initialCapacity int) *StringBuilder {
	return &StringBuilder{data: make([]rune, initialCapacity)}
}

// Creates a new instance of the StringBuilder with a preallocated text
func NewFromString(text string) *StringBuilder {
	return &StringBuilder{
		data:     []rune(text),
		position: len(text),
	}
}

// Appends a text to the StringBuilder instance
func (s *StringBuilder) Append(text string) {
	newLen := s.position + len(text)
	if newLen > cap(s.data) {
		s.grow(newLen)
	}

	copy(s.data[s.position:], []rune(text))
	s.position = newLen
}

// Appends a text and a new line character to the StringBuilder instance
func (s *StringBuilder) AppendLine(text string) {
	s.Append(text)
	s.Append("\n")
}

// Appends a single character to the StringBuilder instance
func (s *StringBuilder) AppendRune(char rune) {
	newLen := s.position + 1
	if newLen > cap(s.data) {
		s.grow(newLen)
	}

	s.data[s.position] = char
	s.position++
}

// Returns the current length of the represented string
func (s *StringBuilder) Len() int {
	return s.position
}

// Returns the represented string
func (s *StringBuilder) ToString() string {
	return string(s.data[:s.position])
}

func (s *StringBuilder) Remove(start int, length int) error {
	if start >= s.position {
		return fmt.Errorf("start is after the end of the string")
	}
	if start < 0 {
		return fmt.Errorf("start can't be a negative value")
	}
	if length < 0 {
		return fmt.Errorf("length can't be a negative value")
	}

	endIndex := start + length - 1

	if endIndex > s.position {
		return fmt.Errorf("can't delete after the end of the string")
	}

	if length == 0 {
		return nil
	}

	s.data = append(s.data[:start], s.data[endIndex])
	s.position -= length

	return nil
}

func (s *StringBuilder) Insert(index int, text string) error {
	if index < 0 {
		return fmt.Errorf("index can't be negative")
	}

	if index > s.position {
		return fmt.Errorf("can't write outside the buffer")
	}

	newLen := s.position + len(text)
	if newLen >= cap(s.data) {
		s.grow(newLen)
	}

	s.data = append(s.data[:index], append([]rune(text), s.data[index:]...)...)
	s.position = newLen

	return nil
}

func (s *StringBuilder) grow(lenToAdd int) {
	// Grow times 2 until lenToAdd fits
	newLen := cap(s.data)

	if cap(s.data) == 0 {
		newLen = 8
	}

	for newLen < lenToAdd {
		newLen = newLen * 2
	}

	new := make([]rune, newLen)
	copy(new, s.data)
	s.data = new
}
