def sum_of_multiples(l, m):
    return sum(i for i in range(l) if any(i % j == 0 for j in m if j > 0)) if not [] else 0