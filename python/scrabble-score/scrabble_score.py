


styles_column_headers = ['Ballot Style']
styles_column = next(
            map(lambda valid_header: data_frame[valid_header], filter(
                    lambda header: header in data_frame
                    and not data_frame[header].empty,
                    styles_column_headers)), None
            )


styles_column_headers = ['Ballot Style']


def get_style_columns(data_frame):
    return [data_frame.get(header) for header
            in styles_column_headers if data_frame.get(header) if is not None]


def head_option(values):
    return next(iter(values), None)


style_column = head_option(get_style_columns(data_frame))



# def hamming(str_a, str_b):
#     return sum(a != b for (a, b) in zip(str_a, str_b))
#
#
# if __name__ == "__main__":
#     A1 = "No. 11 Constitutional Revision\nArticle |, Section 2, Article X, " \
#             " Section 13, Article XII, New Section Lobbying and Abuse of Office by Public Officers"
#     A2 = "No. 1Z Constitutional Revision\nArticle II, Section 8, Article V," \
#          "Section 9 and 19 Property Rights; Removal of Obsolet Provision; Criminal Statutes"
#     s = time.time()
#     compare_strings(A1, A2)
#     e = time.time()
#     print(f"{e - s}")
#
#     s = time.time()
#     seq = SequenceMatcher(None, A1, A2).ratio()
#     e = time.time()
#     print(f"{e - s}")
#
#     s = time.time()
#     lev = Levenshtein.ratio(A1, A2)
#     e = time.time()
#     print(f"{e - s}")

#     B1 = 'No. 12 Constitutional Revision'
#     C1 = 'No. 1l Constitutional Revision'
#     D1 = 'No. 11 Constitutional Revision'
#
#     B2 = 'No. 12 Constitutional Revision'
#     C2 = 'No. 1 Constitutional Revision'
#
#     part_a = A1.split('\n')[0]
#     part_b = A2.split('\n')[0]
#     print(f"A1: {A1}")
#     print("-" * len(B1))
#     print(f"A2: {A2}")
#     print("-" * len(B1))
#     print(f"Hamming: {hamming(A1, A2)}")
#     print(f"Compare: {compare_strings(A1, A2)}")
#     print(f"SequenceMatcher: {SequenceMatcher(None, A1, A2).ratio()}")
#     print(f"Levenshtein: {Levenshtein.ratio(A1, A2)}")
#     print("-" * len(B1))
#     print(f"A1_part: {part_a}")
#     print("-" * len(B1))
#     print(f"A2_part: {part_b}")
#     print("-" * len(B1))
#     print(f"Hamming: {hamming(part_a, part_b)}")
#     print(f"Compare: {compare_strings(part_a, part_b)}")
#     print(f"SequenceMatcher: {SequenceMatcher(None, part_a, part_b).ratio()}")
#     print(f"Levenshtein: {Levenshtein.ratio(part_a, part_b)}")
