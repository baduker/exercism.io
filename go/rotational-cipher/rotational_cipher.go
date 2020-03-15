package rotationalcipher

import (
	"fmt"
	"strings"
)

type action struct {
	encode func(rune) rune
}

func newCaesar(key int) (*action, bool) {
	if key < 0 || key > 26 {
		return nil, false
	}
	rotKey := rune(key)
	return &action{
		encode: func(ch rune) rune {
			// covers ASCII letters
			if ch >= 'a' && ch <= 'z'-rotKey || ch >= 'A' && ch <= 'Z'-rotKey {
				return ch + rotKey
			// skips anything else that's not a letter
			} else if ch > 'z'-rotKey && ch <= 'z' || ch > 'Z'-rotKey && ch <= 'Z' {
				return ch + rotKey - 26
			}
			return ch
		},
	}, true
}

func (act action) cipher(plainText string) string {
	return strings.Map(act.encode, plainText)
}

func RotationalCipher(plainText string, rotKey int) string {
	newRotCipher, ok := newCaesar(rotKey)
	if !ok {
		fmt.Println("Key:", rotKey, "is invalid.")
	}
	return newRotCipher.cipher(plainText)
}
