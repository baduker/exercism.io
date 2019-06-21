def is_pangram(sentence):
    alpha = [chr(i) for i in range(97, 123)]
    return not (set(''.join(alpha)) - set(''.join(sentence.split()).lower()))