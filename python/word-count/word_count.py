import re
from collections import Counter


def count_words(sentence):
    return Counter(
        re.findall(r"[\w0-9]+'?[\w0-9]*(?<!')", sentence.lower()))
