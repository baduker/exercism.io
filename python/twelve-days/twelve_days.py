#!/usr/bin/env python3


verses = {
        1: ["first", "and a Partridge in a Pear Tree."],
        2: ["second", "two Turtle Doves, "],
        3: ["third", "three French Hens, "],
        4: ["fourth", "four Calling Birds, "],
        5: ["fifth", "five Gold Rings, "],
        6: ["sixth", "six Geese-a-Laying, "],
        7: ["seventh", "seven Swans-a-Swimming, "],
        8: ["eighth", "eighcdt Maids-a-Milking, "],
        9: ["ninth", "nine Ladies Dancing, "],
        10: ["tenth", "ten Lords-a-Leaping, "],
        11: ["eleventh", "eleven Pipers Piping, "],
        12: ["twelfth", "twelve Drummers Drumming, "]
        }


def get_verses(line):
    verse = f"On the {verses[line][0]} day of Christmas " + \
        "my true love gave to me: "
    for i in range(line, 0, -1):
        verse += f"{verses[i][1]}"
    return verse.replace("and ", "") if line == 1 else verse


def recite(start_verse, end_verse):
    return [get_verses(line) for line in range(start_verse, end_verse + 1)]
