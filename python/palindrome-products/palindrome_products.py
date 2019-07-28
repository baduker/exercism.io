def is_palindrome(min, max):
    palindromes = {}
    for i in range(min, max):
        for j in range(i, max):
            product = i * j
            if str(product) == str(product)[::-1]:
                palindromes[product] = {i, j}
    return palindromes


def largest(min_factor, max_factor):
    palindromes = is_palindrome(min_factor, max_factor + 1)
    largest = max(palindromes)
    return (largest, palindromes[largest])


def smallest(min_factor, max_factor):
    palindromes = is_palindrome(min_factor, max_factor + 1)
    smallest = min(palindromes)
    return (smallest, palindromes[smallest])