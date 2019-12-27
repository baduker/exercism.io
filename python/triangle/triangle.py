def is_triangle(sides):
    a, b, c = sorted(sides)
    return a > 0 and a + b > c


@is_triangle
def equilateral(sides: list) -> bool:
    a, b, c = sides
    return a == b == c


@is_triangle
def isosceles(sides: list) -> bool:
    return len(set(sides)) < 3


@is_triangle
def scalene(sides: list) -> bool:
    return len(set(sides)) == 3
