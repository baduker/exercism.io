SOUNDS = {3: "Pling", 5: "Plang", 7: "Plong"}
FACTORS = (3, 5, 7)

def is_divisible_by(number, divisior):
    return number % divisior == 0


def raidndrops(number):
    return [SOUNDS[factor] for factor in FACTORS if is_divisible_by(number, factor)]


def convert(number):
    speak = raidndrops(number)
    return "".join(speak) if speak else str(number)
