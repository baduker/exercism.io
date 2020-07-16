import string


def rotate(text, key):
    alphabet = string.ascii_uppercase + string.ascii_lowercase
    abc = string.ascii_lowercase
    ABC = string.ascii_uppercase
    return text.translate(str.maketrans(alphabet, ABC[key:] + ABC[:key] + abc[key:] + abc[:key]))
