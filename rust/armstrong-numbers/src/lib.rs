pub fn is_armstrong_number(num: u32) -> bool {
    
    let my_str = num.to_string();
    let power = my_str.len();
    let mut sum = 0;

    for ch in my_str.chars() {
        sum += ch
            .to_string()
            .parse::<u32>()
            .unwrap()
            .pow(power as u32);
    }

    sum == num
}
