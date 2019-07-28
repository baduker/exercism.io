def slices(series, length):
    if length <= 0 or length > len(series):
        raise ValueError("Invalid data!")
    else:
        return [''.join(map(str, series[i:i + length])) for i in range(len(series) - length + 1)]
