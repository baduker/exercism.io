package bob

import "strings"
import "unicode"

type BobSays struct {
	bobSays string
}

func BobPondersAResponse(remark string) BobSays {
	return BobSays{strings.TrimSpace(remark)}
}

func (b BobSays) IsQuestion() bool {
	return strings.HasSuffix(b.bobSays, "?")
}

func (b BobSays) IsSilence() bool {
	return b.bobSays == ""
}

func (b BobSays) IsShouting() bool {
	JustLetters := strings.IndexFunc(b.bobSays, unicode.IsLetter) >= 0
	Upper := strings.ToUpper(b.bobSays) == b.bobSays
	return Upper && JustLetters
}

func (b BobSays) IsTired() bool {
	return b.IsShouting() && b.IsQuestion()
}

func Hey(remark string) string {
	bobThinks := BobPondersAResponse(remark)
	switch {
	case bobThinks.IsSilence():
		return "Fine. Be that way!"
	case bobThinks.IsTired():
		return "Calm down, I know what I'm doing!"
	case bobThinks.IsShouting():
		return "Whoa, chill out!"
	case bobThinks.IsQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
