import string
import collections


def normalize_words(sentence):

    punctuation_map = str.maketrans({p: ' ' for p in string.punctuation if p != "'"})
    map_words_from_sentence = []
    sentence = sentence.translate(punctuation_map)
    for word in sentence.split():
        word = word.lower()
        map_words_from_sentence.append((word, 1))
    return map_words_from_sentence 


def map_word_occurance(mapped_words):
    word_occurance = collections.defaultdict(list)
    for word, count in mapped_words:
        word_occurance[word].append(count)
    return word_occurance.items()


def reduce(word_count):
    reduced = dict()
    for pair in word_count:
        word, count = pair
        reduced[word] = sum(count)
    return reduced


def count_words(sentence: str) -> dict:
    # print(normalize_words(sentence))
    # print(reduce(map_word_occurance(normalize_words(sentence))))
    return reduce(map_word_occurance(normalize_words(sentence)))
