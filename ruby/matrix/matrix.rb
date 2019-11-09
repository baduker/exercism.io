class Matrix

  attr_reader :rows, :columns

  def initialize(matrix_from_string)
    @rows = get_rows(matrix_from_string)
    @columns = rows.transpose
  end

  private

  def get_rows(string)
    string.each_line.map { |row| row.split.map(&:to_i) }
  end

end
