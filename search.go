package Text

// Returns all occurrences of needle in haystack
func findAll(haystack []rune, needle string) []int {
	needleAsRunes := []rune(needle)
	lenNeedle := len(needleAsRunes)

	if len(haystack) == 0 || len(needle) == 0 || lenNeedle > len(haystack) {
		return []int{}
	}

	var items []int

	limit := len(haystack) - lenNeedle
	for i := 0; i <= limit; i++ {
		if matchAt(haystack, needleAsRunes, i) {
			items = append(items, i)
		}
	}

	return items
}

// Returns the first occurrence of haystack in needle or -1 if not found.
func findFirst(haystack []rune, needle string) int {
	needleAsRunes := []rune(needle)
	lenNeedle := len(needleAsRunes)

	if len(haystack) == 0 || len(needle) == 0 || lenNeedle > len(haystack) {
		return -1
	}

	limit := len(haystack) - lenNeedle
	for i := 0; i <= limit; i++ {
		if matchAt(haystack, needleAsRunes, i) {
			return i
		}
	}

	return -1
}

// Returns the last occurrence of haystack in needle or -1 if not found.
func findLast(haystack []rune, needle string) int {
	needleAsRunes := []rune(needle)
	lenNeedle := len(needleAsRunes)

	if len(haystack) == 0 || len(needle) == 0 || lenNeedle > len(haystack) {
		return -1
	}

	limit := len(haystack) - lenNeedle
	for i := limit; i >= 0; i-- {
		if matchAt(haystack, needleAsRunes, i) {
			return i
		}
	}

	return -1
}

// Checks if the needle matches within the haystack starting at a given position.
func matchAt(haystack []rune, needle []rune, pos int) bool {
	for j := range needle {
		if haystack[pos+j] != needle[j] {
			return false
		}
	}
	return true
}
