def proteins(strand):
    RNA = {
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    }

    STOP = ["UAA", "UAG", "UGA"]
    codons = list(map(''.join, zip(*[iter(strand)] * 3)))

    translation = []
    for codon in codons:
      if codon in STOP:
        break
      else:
        translation.append(RNA.get(codon))
    return translation