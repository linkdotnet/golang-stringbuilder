package Text

type StringBuilder struct {
	data     []rune
	position int
}

// Creates a new instance of the StringBuilder with preallocated array
func NewStringBuilder(initialCapacity int) *StringBuilder {
	return &StringBuilder{data: make([]rune, initialCapacity)}
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
