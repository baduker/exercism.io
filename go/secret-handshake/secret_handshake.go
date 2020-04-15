package secret


func Handshake(number uint) (secret []string) {
	if number&1 == 1 {
		secret = append(secret, "wink")
	}
	if number&2 == 2 {
		secret = append(secret, "double blink")
	}
	if number&4 == 4 {
		secret = append(secret, "close your eyes")
	}
	if number&8 == 8 {
		secret = append(secret, "jump")
	}
	if number&16 == 16 {
		for i, j := 0, len(secret) - 1; i < len(secret) / 2; i, j = i + 1, j - 1 {
			secret[i], secret[j] = secret[j], secret[i]
		}
	}
	return
}
