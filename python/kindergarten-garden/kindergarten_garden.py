class Garden:
    students = [
        'Alice', 'Bob', 'Charlie', 'David', 'Eve', 'Fred',
        'Ginny', 'Harriet', 'Ileana', 'Joseph', 'Kincaid',
        'Larry',
        ]

    flower_map = {
        'C': 'Clover',
        'G': 'Grass',
        'R': 'Radishes',
        'V': 'Violets',
        }

    def __init__(self, diagram, students=students):
        self.rows = diagram.splitlines()
        self.students = sorted(students)
        self.student_index = {name:index for (index, name) in enumerate(self.students)}

    def plants(self, student):
        plant_row = self.student_index[student] * 2
        letters = ''.join(row[plant_row : plant_row + 2] for row in self.rows)
        return [self.flower_map[letter] for letter in letters]
