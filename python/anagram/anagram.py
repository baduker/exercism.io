def find_anagrams(word, candidates):
    return [sample for sample in candidates if is_anagram(word, sample)]


def is_anagram(word_one, word_two):
    word_one = word_one.lower()
    word_two = word_two.lower()

    same_letters = sorted(word_one) == sorted(word_two)
    same_word = word_one != word_two

    return same_letters and same_word