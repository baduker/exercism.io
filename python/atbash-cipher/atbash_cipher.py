import string


CIPHER = {
    k:v for k, v in zip(string.ascii_letters, string.ascii_letters[::-1])
    }


def encode(plain_text, chunk_size=5):
    sanitaize = plain_text.translate(str.maketrans("", "", string.punctuation))
    encoded = [CIPHER.get(l, l) for l in sanitaize.replace(" ", "")]
    chunked = [
    ''.join(encoded[i:i + chunk_size]).lower() for i 
    in range(0, len(encoded), chunk_size)
    ]
    return ' '.join(chunked)


def decode(ciphered_text):
    reverse_cipher = {v:k for k, v in CIPHER.items()}
    return "".join(
        reverse_cipher.get(l, l) for l in ciphered_text.replace(" ", "")
        ).lower()
