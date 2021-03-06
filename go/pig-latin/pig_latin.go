package piglatin

eximport "strings"

const pigTail = "ay"

func PigIt(input string) string {
	if len(input) == 0 {
		return ""
	}

	words := strings.Fields(strings.ToLower(input))
	result := make([]string, 0, len(words))
	for _, word := range words {
		if len(word) < 3 {
			result = append(result, piggify(word, 1))
			continue
		}
		switch word[:3] {
		case "squ", "sch", "thr":
			result = append(result, piggify(word, 3))
			continue
		}
		switch word[:2] {
		case "ch", "qu", "th", "rh":
			result = append(result, piggify(word, 2))
			continue
		case "ye", "xe":
			result = append(result, piggify(word, 1))
			continue
		}
		switch word[:1] {
		case "a", "e", "o", "u", "i", "y", "x":
			result = append(result, word+pigTail)
			continue
		}
		result = append(result, piggify(word, 1))
	}
	return strings.Join(result, " ")
}

func piggify(word string, index int) string {
	return word[index:] + word[:index] + pigTail
}
