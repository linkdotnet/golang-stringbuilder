package Text

// Returns all occurences of needle in haystack
func findAll(haystack []rune, needle string) []int {

	if len(haystack) == 0 || len(needle) == 0 || len(needle) > len(haystack) {
		return []int{}
	}

	items := make([]int, 0, 8)

	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != rune(needle[j]) {
				break
			}

			if j == len(needle)-1 {
				items = append(items, i)
			}
		}
	}

	return items
}

// Returns the first occurrence of haystack in needle or -1 if not found.
func findFirst(haystack []rune, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 || len(needle) > len(haystack) {
		return -1
	}

	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		for j := 0; j < lenNeedle; j++ {
			if haystack[i+j] != rune(needle[j]) {
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
	if len(haystack) == 0 || len(needle) == 0 || len(needle) > len(haystack) {
		return -1
	}

	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	for i := lenHaystack - lenNeedle + 1; i >= 0; i-- {
		for j := 0; j < lenNeedle; j++ {
			if haystack[i+j] != rune(needle[j]) {
				break
			}

			if j == lenNeedle-1 {
				return i
			}
		}
	}

	return -1
}
