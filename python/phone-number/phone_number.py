import re

PHONE_REG_PAT = r"1?\D*(\d{3})\D*(\d{3})\D*(\d{4})\D*$"

class Phone(object):
    def __init__(self, phone_number):
      match_phone_number = re.search(PHONE_REG_PAT, phone_number)
      self.number = "".join(match_phone_number.groups())

    def pretty(number):
      return "({}) {}-{}".format(self.number[:3], self.number[3:6],
              self.number[6:])
    