def is_triangle(sides):
    if min(sides) <= 0:
        return False
    if sum(sorted(sides)[:-1]) < sorted(sides)[-1]:
        return False
    return True


def equilateral(sides):
    triangle = is_triangle(sides)
    if triangle:
        x, y, z = sides
        return x == y == z
    else:
        return False


def isosceles(sides):
    triangle = is_triangle(sides)
    if triangle:
        x, y, z = sides
        return x == y or y == z or z == x
    else:
        return False


def scalene(sides):
    if equilateral(sides) or isosceles(sides):
        return False
    return is_triangle(sides)