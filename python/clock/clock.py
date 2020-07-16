class Clock:
    def __init__(self, hour : int, minute : int):
        part, self.minute = divmod(minute, 60)
        _, self.hour = divmod(hour + part, 24)

    def __repr__(self) -> str:
        return "%02d:%02d" % (self.hour, self.minute)

    def __eq__(self, other):
        return repr(self) == other

    def __add__(self, minutes):
        return Clock(self.hour, self.minute + minutes)

    def __sub__(self, minutes):
        return Clock(self.hour, self.minute - minutes)
