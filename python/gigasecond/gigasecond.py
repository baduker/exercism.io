from datetime import datetime, timedelta


def add(a_date):
  GIGASECOND = timedelta(seconds = 10**9)
  return a_date + GIGASECOND