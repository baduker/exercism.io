def factors(value):
    number = 2
    factors = []
    while number <= value:
        if value % number == 0:
            factors.append(number)
            value /= number
        else:
            number += 1
    return factors