def palindromes(product_range):
    return (int(x) for x in map(str, product_range) if x == x[::-1])


def find_palindromes(min_factor, max_factor, *, reverse=False):
    if max_factor < min_factor:
        raise ValueError("Invalid input!")

    number_ragne = range(min_factor, max_factor + 1)
    product_range = range(min_factor ** 2, max_factor ** 2 + 1)

    if reverse:
        product_range = reversed(product_range)

    for palindrome in palindromes(product_range):
        factors = [
            [item, palindrome // item]
            for item in number_ragne
            if palindrome % item == 0
            and palindrome // item in number_ragne
            and item <= palindrome // item
        ]
        if factors:
            return (palindrome, factors)

    return None, []


def largest(min_factor, max_factor):
    return find_palindromes(min_factor, max_factor, reverse=True)


def smallest(min_factor, max_factor):
    return find_palindromes(min_factor, max_factor)
