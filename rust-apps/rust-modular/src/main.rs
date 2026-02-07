use bcrypt::{DEFAULT_COST, hash};

fn main() {
    let hashed = hash_password("MyPassword".to_string());
    println!("{}", hashed);
}

fn hash_password(pass: String) -> String {
    match hash(pass, DEFAULT_COST) {
        Ok(hashed) => hashed,
        Err(e) => panic!("error: {}", e),
    }
}
