def distance(str_a, str_b):
    if len(str_a) != len(str_b):
        raise ValueError("Strand lengths are not equal!")
    else:
        return sum(1 for (a, b) in zip(str_a, str_b) if a != b)
