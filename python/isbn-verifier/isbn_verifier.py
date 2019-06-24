import re     

def prase_isbn(isbn):
    isbn_pattern = r"^(\d{9})(\d|X)$"
    match = re.search(isbn_pattern, isbn)
    return match


def remove_dashes(isbn):
    no_dashes = isbn.upper().replace("-", "");
    return no_dashes


def verify_check_digit(parsed_isbn):
    return 10 if parsed_isbn.group(2) == "X" else int(parsed_isbn.group(2))


def isbn_formula(isbn_digits):
    return sum((i + 1) * int(digit) for i, digit in enumerate(isbn_digits))


def is_valid(isbn):
    clean_isbn = remove_dashes(isbn)
    match = prase_isbn(clean_isbn)


    if not match:
        return False

    get_digits = match.group(1)
    check_digit = verify_check_digit(match)

    result = isbn_formula(get_digits) % 11
    return result == check_digit