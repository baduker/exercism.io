import unittest
from datetime import datetime

from gigasecond import add


# Tests adapted from `problem-specifications//canonical-data.json` @ v2.0.0

class GigasecondTest(unittest.TestCase):
    def test_date_only_specification_of_time(self):
        self.assertEqual(
            add(datetime(2011, 4, 25)),
            datetime(2043, 1, 1, 1, 46, 40))

    def test_another_date_only_specification_of_time(self):
        self.assertEqual(
            add(datetime(1977, 6, 13)),
            datetime(2009, 2, 19, 1, 46, 40))

    def test_one_more_date_only_specification_of_time(self):
        self.assertEqual(
            add(datetime(1959, 7, 19)),
            datetime(1991, 3, 27, 1, 46, 40))

    def test_full_time_specified(self):
        self.assertEqual(
            add(datetime(2015, 1, 24, 22, 0, 0)),
            datetime(2046, 10, 2, 23, 46, 40))

    def test_full_time_with_day_roll_over(self):
        self.assertEqual(
            add(datetime(2015, 1, 24, 23, 59, 59)),
            datetime(2046, 10, 3, 1, 46, 39))
    """   
    def test_yourself(self):
        # customize this to test your birthday and find your gigasecond date:
        your_birthday = datetime(1985, 1, 9)
        your_gigasecond = datetime(2301, 11, 29, 17, 46, 40)

        self.assertEqual(add(your_birthday), your_gigasecond)
    """

if __name__ == '__main__':
    unittest.main()
