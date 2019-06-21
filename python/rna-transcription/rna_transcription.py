def to_rna(dna_strand):
    d = {"A": "U", "T": "A", "G": "C", "C": "G"}
    return ''.join(d.get(nucleotide) for nucleotide in dna_strand)
