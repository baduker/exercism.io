#!/usr/bin/python3

def check(number):
    if number <= 0:
      raise ValueError("Error: expected positive number!")
    else:
      return number


def is_even(number):
    return number // 2


def is_odd(number):
    return number * 3 + 1


def steps(number):
    count_steps = 0
    valid_number = check(number)
    while valid_number > 1:
      if valid_number % 2 == 0:
        valid_number = is_even(valid_number)
      else:
        valid_number = is_odd(valid_number)
      count_steps += 1
    return count_steps