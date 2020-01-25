package atbash

import "bytes"

func Atbash(plainText string) string {
	// Calculate the length of an encoded string taking into consideration that
	// output has to be in chucks of length 5
	length := len(plainText) + len(plainText)/5
	// Make a new byte slice to accommodate the encoded bytes
	encoded := make([]byte, 0, length)
	for _, b := range []byte(plainText) {
		switch {
		// 'z' - b (byte value) + 'a' is an algorithm to get the transposed letter of the alphabet
		case b >= 'a' && b <= 'z':
			encoded = append(encoded, 'z'-b+'a')
		// handles uppercase letters => lowercase
		case b >= 'A' && b <= 'Z':
			encoded = append(encoded, 'z'-b+'A')
		// if there's a digit, add it as it is
		case b >= '0' && b <= '9':
			encoded = append(encoded, b)
		default:
			continue
		}
		// Every five bytes inserts a space
		if len(encoded)%6 == 5 {
			encoded = append(encoded, ' ')
		}
	}
	// removes the trailing white space
	return string(bytes.TrimSuffix(encoded, []byte{' '}))
}
