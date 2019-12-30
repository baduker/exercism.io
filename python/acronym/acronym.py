import string


def abbreviate(words):
    punc_map = str.maketrans({
        p: ' ' for p in string.punctuation if p not in "'"
        })
    return "".join(w[0].upper() for w in words.translate(punc_map).split())
