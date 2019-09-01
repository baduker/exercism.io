#!/usr/bin/env python3


verses = [["first", "and a Partridge in a Pear Tree."],
          ["second", "two Turtle Doves, "],
          ["third", "three French Hens, "],
          ["fourth", "four Calling Birds, "],
          ["fifth", "five Gold Rings, "],
          ["sixth", "six Geese-a-Laying, "],
          ["seventh", "seven Swans-a-Swimming, "],
          ["eighth", "eight Maids-a-Milking, "],
          ["ninth", "nine Ladies Dancing, "],
          ["tenth", "ten Lords-a-Leaping, "],
          ["eleventh", "eleven Pipers Piping, "],
          ["twelfth", "twelve Drummers Drumming, "]]


def get_verses(line):
    verse = f"On the {verses[line][0]} day of Christmas " \
    "my true love gave to me: "
    for i in range(line, -1, -1):
        verse += f"{verses[i][1]}"
    return verse.replace("and ", "") if line == 0 else verse


def recite(start_verse, end_verse):
    return [get_verses(line_n)
            for line_n in range(start_verse - 1, end_verse)]