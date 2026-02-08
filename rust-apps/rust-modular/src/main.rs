mod utils;

use utils::hashlib::{hash_password, verify_password};

fn main() {
    let my_password = "HelloMyPassword";
    let hashed = match hash_password(my_password.to_string()) {
        Ok(hashed) => hashed,
        Err(e) => {
            eprintln!("error in hash pass: {e}");
            return;
        }
    };

    println!("hash pass {hashed}");

    match verify_password(my_password, hashed) {
        Ok(is_valid) => {
            if is_valid {
                println!("password is match!");
            } else {
                println!("pass is not match!!");
            }
        }
        Err(e) => {
            println!("error in validate pass: {e}");
            return;
        }
    };
}
