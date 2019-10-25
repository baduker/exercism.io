class Allergies(object):

    ALLERGY_KEYS = ("cats", "pollen", "chocolate", "tomatoes", "strawberries",
                    "shellfish", "peanuts", "eggs")

    def __init__(self, score):
        allergy_values = [bool(int(x)) for x in format(score, '08b')][-8:]
        self.allergies = dict(zip(self.ALLERGY_KEYS, allergy_values))

    def allergic_to(self, item):
        return self.allergies[item]

    @property
    def lst(self):
        return [allergy for (allergy, hasAllergy) in
                self.allergies.items() if hasAllergy]
