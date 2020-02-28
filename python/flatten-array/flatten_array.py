import itertools
import collections

# No recursion
def helper(iterable, any_iter=collections.abc.Iterable):
    remainder = iter(iterable)
    try:
        while True:
            first = next(remainder)
            if isinstance(first, any_iter):
                remainder = itertools.chain(first, remainder)
            else:
                yield first
    except StopIteration:
        return


def flatten(iterable):
    return list(i for i in filter(None.__ne__, helper(iterable)))    
