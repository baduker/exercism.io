from functools import reduce


def get_sum_of_factors(n):
    if n <= 0:
        raise ValueError("Value can't be zero or lower.")

    return sum(list(set(reduce(list.__add__, ([i, n//i] for i in \
        range (1, int(n**0.5) + 1) if n % i == 0))))) - n


def classify(n):
    answers = {1: "perfect", 2: "abundant", 3: "deficient"}

    if n == get_sum_of_factors(n):
        return answers[1]
    elif n < get_sum_of_factors(n):
        return answers[2]
    else:
        return answers[3]