class Luhn:
    def __init__(self, card_num):
        self.card_num = card_num

    def valid(self):
        # sanitize the string
        card_number = self.card_num.strip().replace(" ", "")
        # anything shorter than 2 chars is invalid; bail out
        if len(card_number) < 2:
            return False

        # check if all are digits
        if not all(d.isdigit() for d in card_number):
            return False

        # reverse and convert to integers
        digits = [int(d) for d in card_number[::-1]]
        return (
            sum(digits[0::2]) + # sum of evens
            sum(sum(divmod(d * 2, 10)) for d in digits[1::2]) # get odds
            ) % 10 == 0
