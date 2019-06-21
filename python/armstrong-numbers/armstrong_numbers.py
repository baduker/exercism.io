def is_armstrong_number(number):
    number_to_str = str(number)
    return sum([int(num) ** len(number_to_str) for num in number_to_str]) == number
