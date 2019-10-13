def palindromes(product_range):
    # returns a list of integer palindromes from input
    return (int(x) for x in map(str, product_range) if x == x[::-1])


def find_palindromes(min_factor, max_factor, *, reverse=False):
    # checks for valid input
    if max_factor < min_factor:
        raise ValueError("Invalid input!")

    # build a list of min_factor & max_factor, both inclusive
    number_ragne = range(min_factor, max_factor + 1)
    # build a list of all possible products within the above range
    product_range = range(pow(min_factor, 2), pow(max_factor, 2) + 1)

    # reverse the list of all possible products -> highest first
    if reverse:
        product_range = reversed(product_range)

    # find palindrome and its factor
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
