def is_triangle(func):
    """
    The Triangle Inequality Theorem states: c < a + b. We can simply add c to both sides which then makes it 2*c < a + b + c. This formula only needs to be performed once (as opposed to once for every side) if c is the side with the largest length so it can be simplified to 2*max(s) < sum(s).
    """
    return lambda s: all(s) and 2 * max(s) < sum(s) and func(s)

@is_triangle
def equilateral(sides):
    return len(set(sides)) == 1

@is_triangle
def isosceles(sides):
    return len(set(sides)) < 3

@is_triangle
def scalene(sides):
    return len(set(sides)) == 3