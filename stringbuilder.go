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
func NewStringBuilderFromString(text string) *StringBuilder {
	textRunes := []rune(text)
	return &StringBuilder{
		data:     textRunes,
		position: len(textRunes),
	}
}

// Appends a text to the StringBuilder instance
func (s *StringBuilder) Append(text string) {
	textRunes := []rune(text)
	newLen := s.position + len(textRunes)
	if newLen > cap(s.data) {
		s.grow(newLen)
	}

	copy(s.data[s.position:], textRunes)
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

	x := start + length
	copy(s.data[start:], s.data[x:])
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

	runeText := []rune(text)
	newLen := s.position + len(runeText)
	if newLen >= cap(s.data) {
		s.grow(newLen)
	}

	s.data = append(s.data[:index], append(runeText, s.data[index:]...)...)
	s.position = newLen

	return nil
}

// Removes all characters from the current instance. This sets the internal size to 0.
// The internal array will stay the same.
func (s *StringBuilder) Clear() {
	s.position = 0
}

// Gets the rune at the specific position
func (s *StringBuilder) RuneAt(index int) rune {
	return s.data[index]
}

// Returns the string builder as a rune-slice. Be careful as this returns the internal slice.
// Changes to that will reflect in this string builder instance.
func (s *StringBuilder) AsRuneSlice() []rune {
	return s.data[:s.position]
}

// Returns the first occurrence of the given text in the string builder. Returns -1 if not found
func (s *StringBuilder) FindFirst(text string) int {
	return findFirst(s.AsRuneSlice(), text)
}

// Returns the last occurrence of the given text in the string builder. Returns -1 if not found
func (s *StringBuilder) FindLast(text string) int {
	return findLast(s.AsRuneSlice(), text)
}

// Returns all occurrences of the given text in the string builder. Returns an empty if no occurrence found.
func (s *StringBuilder) FindAll(text string) []int {
	return findAll(s.AsRuneSlice(), text)
}

// Replaces all occurrences of oldValue with newValue
func (s *StringBuilder) ReplaceRune(oldValue rune, newValue rune) {
	occurrences := s.FindAll(string(oldValue))

	for _, v := range occurrences {
		s.data[v] = newValue
	}
}

// Replaces all occurrences of oldValue with newValue
func (s *StringBuilder) Replace(oldValue string, newValue string) {
	if oldValue == newValue {
		return
	}

	occurrences := s.FindAll(oldValue)

	oldValueRunes := []rune(oldValue)
	newValueRunes := []rune(newValue)

	delta := len(newValueRunes) - len(oldValueRunes)

	for i := 0; i < len(occurrences); i++ {
		index := occurrences[i] + delta*i

		// newValue is smaller than old value
		// We can insert the slice and remove the overhead
		if delta < 0 {
			copy(s.data[index:], newValueRunes)
			s.Remove(index+len(newValueRunes), -delta)
		} else if delta == 0 {
			// Same length -> We can just replace the memory slice
			copy(s.data[index:], newValueRunes)
		} else {
			// newValue is larger than the old value
			// First add until the old memory region
			// and insert afterwards the rest
			oldLen := len(oldValueRunes)
			copy(s.data[index:], []rune(newValueRunes[:oldLen]))
			s.Insert(index+oldLen, string(newValueRunes[len(oldValueRunes):]))
		}
	}
}

// Implements the io.Writer interface so the StringBuilder can be used with fmt.Printf
func (s *StringBuilder) Write(p []byte) (int, error) {
	before := s.Len()
	s.Append(string(p))
	delta := s.Len() - before

	return delta, nil
}

func (s *StringBuilder) grow(lenToAdd int) {
	// Grow times 2 until lenToAdd fits
	newLen := len(s.data)

	if newLen == 0 {
		newLen = 8
	}

	for newLen < lenToAdd {
		newLen = newLen * 2
	}

	new := make([]rune, newLen)
	copy(new, s.data)
	s.data = new
}
