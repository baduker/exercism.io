def find_primes(n):
    multiples = set()
    for number in range(2, n+1):
        if number not in multiples:
            yield number
            multiples.update(range(number*number, n+1, number))


def primes(limit):
    return list(find_primes(limit))
