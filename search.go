package Text

var arrayEmpty = make([]int, 0)

func findAllText(text []rune, needle string) []int {
	if len(text) == 0 || len(needle) == 0 || len(text) < len(needle) {
		return arrayEmpty
	}

	return arrayEmpty
}
