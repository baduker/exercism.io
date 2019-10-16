import re
from collections import Counter

def count_words(sentence):
    # with a negative lookbehind the regex grabs all words with ' & " quotes
    sanitization_patter = r"([A-Za-z0-9]+'?[A-Za-z0-9]*)(?<!')"
    sanitized_words = re.findall(sanitization_patter, sentence.lower())
    return Counter(sanitized_words)
