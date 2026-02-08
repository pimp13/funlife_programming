fn zarib_3_ya_5(n: u32) -> bool {
    if n % 3 == 0 || n % 5 == 0 {
        true
    } else {
        false
    }
}

fn main() {
    let mut sum: u32 = 0;
    for i in 1..1000 {
        if zarib_3_ya_5(i) {
            sum += i;
        }
    }
    println!("{sum}");
}
