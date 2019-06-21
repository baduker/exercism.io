from itertools import dropwhile
from random import choice


class Robot:


    used_robots = set()


    def __init__(self):
      self.reset()


    @staticmethod
    def create_random_name():

      letters = ''.join(chr(i) for i in range(65, 91))
      digits = ''.join(str(j) for j in range(0, 11))

      while True:
        random_name = [letters] * 2 + [digits] * 3
        yield ''.join(choice(part_of_name) for part_of_name in random_name)


    def get_random_name(self):
      def used_robot_names(name): return name in self.used_robots
      robot_names = self.create_random_name()
      return next(dropwhile(used_robot_names, robot_names))


    def reset(self):
      self.name = self.get_random_name()
      self.used_robots.add(self.name)