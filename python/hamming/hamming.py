def distance(str_a, str_b):
    if len(str_a) != len(str_b):
        raise ValueError("Strand lengths are not equal!")
    else:
        return sum(str_a != str_b for str_a, str_b in zip(str_a, str_b))
