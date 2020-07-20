pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    (1..limit)
    .filter(|i| factors
        .iter()
        .any(|x| x > &0 && i % x == 0)
    )
    .sum()
}
