SOUNDS = [(3, "Pling"), (5, "Plang"), (7, "Plong")]


def raindrops(number):
    return [sound for d, sound in SOUNDS if number % d == 0]


def convert(number):
    return "".join(raindrops(number)) or str(number)