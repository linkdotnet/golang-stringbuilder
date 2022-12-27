package Text

// Returns all occurences of needle in haystack
func findAll(haystack []rune, needle string) []int {

	needleAsRunes := []rune(needle)
	lenNeedle := len(needleAsRunes)

	if len(haystack) == 0 || len(needle) == 0 || lenNeedle > len(haystack) {
		return []int{}
	}

	items := make([]int, 0, 8)

	for i := 0; i < len(haystack); i++ {
		for j := 0; j < lenNeedle; j++ {
			if haystack[i+j] != needleAsRunes[j] {
				break
			}

			if j == lenNeedle-1 {
				items = append(items, i)
			}
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

	lenHaystack := len(haystack)

	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		for j := 0; j < lenNeedle; j++ {
			if haystack[i+j] != needleAsRunes[j] {
				break
			}

			if j == lenNeedle-1 {
				return i
			}
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

	lenHaystack := len(haystack)

	for i := lenHaystack - lenNeedle; i >= 0; i-- {
		for j := 0; j < lenNeedle; j++ {
			if haystack[i+j] != needleAsRunes[j] {
				break
			}

			if j == lenNeedle-1 {
				return i
			}
		}
	}

	return -1
}
