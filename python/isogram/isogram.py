def is_isogram(s):
  if isinstance(s, str):
      s = [i for i in s.lower() if i.isalpha()]
      return len(s) == len(set(s))
  else:
      raise TypeError("This doesn't look like a string.")