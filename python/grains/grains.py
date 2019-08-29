def validate(number):
    if not 0 < number <= 64:
        raise ValueError("Invalid input!")

def square(number):
    validate(number)
    return 2 ** (number - 1)


def total(number):
    validate(number)
    return sum(square(i + 1) for i in range(number))
