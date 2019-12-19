package acronym

import "strings"

// Abbreviate converts a phrase to its acronym
func Abbreviate(text string) string {
	cleanText := strings.Fields(Replacer(text))
	var acronym []string
	for _, word := range cleanText {
		acronym = append(acronym, strings.ToUpper(word[:1]))
	}
	return strings.Join(acronym, "")
}

func Replacer(s string) string {
	ignored := [2]string{"-", "_"}
	for _, item := range ignored {
		s = strings.Replace(s, item, " ", -1)
	}
	return s
}
