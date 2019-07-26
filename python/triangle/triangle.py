def is_triangle(a, b, c):
    a, b, c = storted([a, b, c])
    return a > 0 and a + b > c


@is_triangle
def is_equilateral(a, b, c):
    return a == b == c


@is_triangle
def is_isosceles(a, b, c):
    return len(set([a, b, c])) < 3


@is_triangle
def is_scalene(a, b, c):
    return len(set([a, b, c])) == 3