pub fn raindrops(n: u32) -> String {
    let mut s = Vec::new();

    if n % 3 == 0 {
        s.push("Pling")
    } else if n % 5 == 0 {
        s.push("Plang")
    } else if n % 7 == 0 {
        s.push("Plong")
    } else {
        return n.to_string();
    }

    s.join("")
}
