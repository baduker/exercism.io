import re
from collections import Counter


def count_words(sentence):
    return Counter(
        re.findall(r"([A-Za-z0-9]+'?[A-Za-z0-9]*)(?<!')", sentence.lower()))
