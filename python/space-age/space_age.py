from typing import Callable, Optional


SECONDS_IN_EARTH_YEAR = 31557600
ORBITAL_PERIODS = {
    'earth': 1,
    'mercury': 0.2408467,
    'venus': 0.61519726,
    'mars': 1.8808158,
    'jupiter': 11.862615,
    'saturn': 29.447498,
    'uranus': 84.016846,
    'neptune': 164.79132,
}


def age_on_planet(seconds, planet):
  annual_seconds = SECONDS_IN_EARTH_YEAR * ORBITAL_PERIODS[planet]
  return round(seconds / annual_seconds, 2)


class SpaceAge(object):


    def __init__(self, seconds):
        self.seconds = seconds
        

    def __getattr__(self, planet):
      if planet.startswith("on_"):
        return lambda: age_on_planet(self.seconds, planet[3:])
      return None