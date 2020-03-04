class Clock:
    def __init__(self, hour, minute):
        self.minutes = (minute + 60 * hour) % 1440

    def __repr__(self):
        return f"{self.minutes // 60:02d}:{self.minutes % 60:02d}"

    def __eq__(self, other):
        return repr(self) == repr(other)

    def __add__(self, minutes):
        self.minutes = (self.minutes + minutes) % 1440
        return self

    def __sub__(self, minutes):
        return self.__add__(-minutes)
