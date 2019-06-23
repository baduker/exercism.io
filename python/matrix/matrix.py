def get_matrix_from_string(s):
  return [[int(x) for x in row.split()] for row in s.splitlines()]


def transpose(matrix):
  return [list(column) for column in zip(*matrix)]


class Matrix(object):


    def __init__(self, matrix_string):
      create_matrix = get_matrix_from_string(matrix_string)
      transpose_matrix = transpose(create_matrix)
      self.rows, self.columns = create_matrix, transpose_matrix


    def row(self, index):
        return self.rows[index-1]


    def column(self, index):
        return self.columns[index-1]