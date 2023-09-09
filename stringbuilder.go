package Text

import (
	"fmt"
	"io"
	"strconv"
	"sync"
)

// StringBuilder is a thread-safe & mutable sequence of characters
type StringBuilder struct {
	data     []rune
	position int
	mu       sync.Mutex
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
		mu:       sync.Mutex{},
	}
}

// Appends a text to the StringBuilder instance
func (s *StringBuilder) Append(text string) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.resize(text)
	textRunes := []rune(text)
	copy(s.data[s.position:], textRunes)
	s.position = s.position + len(textRunes)

	return s
}

// Appends a text and a new line character to the StringBuilder instance
func (s *StringBuilder) AppendLine(text string) *StringBuilder {
	return s.Append(text + "\n")
}

// Appends a single character to the StringBuilder instance
func (s *StringBuilder) AppendRune(char rune) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	newLen := s.position + 1
	if newLen >= cap(s.data) {
		s.grow(newLen)
	}
	s.data[s.position] = char
	s.position++

	return s
}

// Appends a single integer to the StringBuilder instance
func (s *StringBuilder) AppendInt(integer int) *StringBuilder {
	return s.Append(strconv.Itoa(integer))
}

// Appends a single boolean to the StringBuilder instance
func (s *StringBuilder) AppendBool(flag bool) *StringBuilder {
	return s.Append(strconv.FormatBool(flag))
}

// Appends a list of strings to the StringBuilder instance
func (s *StringBuilder) AppendList(words []string) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.resize(words...)
	for _, word := range words {
		textRunes := []rune(word)
		copy(s.data[s.position:], textRunes)
		s.position = s.position + len(textRunes)
	}
	return s
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
	s.mu.Lock()
	defer s.mu.Unlock()

	return remove(start, length, s)
}

func (s *StringBuilder) Insert(index int, text string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return insert(index, text, s)
}

// Removes all characters from the current instance. This sets the internal size to 0.
// The internal array will stay the same.
func (s *StringBuilder) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

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
func (s *StringBuilder) ReplaceRune(oldValue rune, newValue rune) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	occurrences := s.FindAll(string(oldValue))

	for _, v := range occurrences {
		s.data[v] = newValue
	}

	return s
}

// Replaces all occurrences of oldValue with newValue
func (s *StringBuilder) Replace(oldValue string, newValue string) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	if oldValue == newValue {
		return s
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
			remove(index+len(newValueRunes), -delta, s)
		} else if delta == 0 {
			// Same length -> We can just replace the memory slice
			copy(s.data[index:], newValueRunes)
		} else {
			// newValue is larger than the old value
			// First add until the old memory region
			// and insert afterwards the rest
			oldLen := len(oldValueRunes)
			copy(s.data[index:], []rune(newValueRunes[:oldLen]))
			insert(index+oldLen, string(newValueRunes[len(oldValueRunes):]), s)
		}
	}

	return s
}

// Implements the io.Writer interface so the StringBuilder can be used with fmt.Printf
func (s *StringBuilder) Write(p []byte) (int, error) {
	before := s.Len()
	s.Append(string(p))
	delta := s.Len() - before

	return delta, nil
}

// Trims the given characters from the start and end of the string builder or all whitespaces if no characters are given
func (s *StringBuilder) Trim(chars ...rune) *StringBuilder {
	return s.TrimStart(chars...).TrimEnd(chars...)
}

// Trims the given characters from the start of the string builder or all whitespaces if no characters are given
func (s *StringBuilder) TrimStart(chars ...rune) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	start := 0
	trimSet := createTrimSet(chars...)

	for _, ch := range s.data[:s.position] {
		if _, exists := trimSet[ch]; !exists {
			break
		}
		start++
	}

	if start > 0 {
		copy(s.data, s.data[start:s.position])
		s.position -= start
	}

	return s
}

// Trims the given characters from the start of the string builder or all whitespaces if no characters are given
func (s *StringBuilder) TrimEnd(chars ...rune) *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	end := s.position
	trimSet := createTrimSet(chars...)

	for i := s.position - 1; i >= 0; i-- {
		if _, exists := trimSet[s.data[i]]; !exists {
			break
		}
		end--
	}

	s.position = end

	return s
}

// Returns the internal array of the string builder. Be careful as this returns the internal slice.
// Changes to that will reflect in this string builder instance.
func (s *StringBuilder) AsRuneArray() []rune {
	return s.data
}

// Reverses the characters of a string builder
func (s *StringBuilder) Reverse() *StringBuilder {
	s.mu.Lock()
	defer s.mu.Unlock()

	for left, right := 0, s.position-1; left < right; left, right = left+1, right-1 {
		s.data[left], s.data[right] = s.data[right], s.data[left]
	}

	return s
}

// Returns a substring from start (inclusive) to end (exclusive).
func (s *StringBuilder) Substring(start, end int) (string, error) {
	if start < 0 {
		return "", fmt.Errorf("start should always be greater than or equal to zero")
	}
	if end > s.position {
		return "", fmt.Errorf("end cannot be greater than the length of string builder")
	}
	if start > end {
		return "", fmt.Errorf("start cannot be greater than the end for Substring() function")
	}
	r := make([]rune, end-start)
	copy(r, s.data[start:end])
	return string(r), nil
}

// Reads the content from the given reader and appends it to the string builder
func (s *StringBuilder) ReadFrom(r io.Reader) (int64, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	text := string(bytes)
	s.Append(text)
	return int64(len(text)), nil
}

func (s *StringBuilder) resize(words ...string) {
	allWordLength := 0
	for _, word := range words {
		allWordLength += len(word)
	}
	newLen := s.position + allWordLength
	if newLen > cap(s.data) {
		s.grow(newLen)
	}
}

func remove(start int, length int, s *StringBuilder) error {
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

func insert(index int, text string, s *StringBuilder) error {
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

func (s *StringBuilder) grow(lenToAdd int) {
	// Grow times 2 until lenToAdd fits
	newLen := len(s.data)

	if newLen == 0 {
		newLen = 8
	}

	for newLen < lenToAdd {
		newLen = newLen * 2
	}

	s.data = append(s.data, make([]rune, newLen-len(s.data))...)
}

func createTrimSet(chars ...rune) map[rune]bool {
	trimSet := make(map[rune]bool)

	if len(chars) == 0 {
		trimSet[' '] = true
		trimSet['\t'] = true
		trimSet['\n'] = true
		trimSet['\r'] = true
	} else {
		for _, ch := range chars {
			trimSet[ch] = true
		}
	}

	return trimSet
}
