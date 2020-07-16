def transform(ld):
    return {w.lower(): k for k, v in ld.items() for w in v}
