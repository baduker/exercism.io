from functools import wraps


def check_if_number(func):
    @wraps(func)
    def wrapper(number):
        return is_number(number) and func(number)
    return wrapper

    
def is_number(number):
    if number <= 0:
        raise ValueError("Error: expected a positive number!")
    else:
        return number


@check_if_number
def is_even(number):
    return number // 2


@check_if_number
def is_odd(number):
    return number * 3 + 1


@check_if_number
def steps(number):
    count_steps = 0
    while number > 1:
      if number % 2 == 0:
        number = is_even(number)
      else:
        number = is_odd(number)
      count_steps += 1
    return count_steps