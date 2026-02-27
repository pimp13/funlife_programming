use std::{thread, time::Duration};

fn main() {
    let handler = thread::spawn(|| {
        for i in 1..=100 {
            println!("spawned thread says: {}", i);
            thread::sleep(Duration::from_secs(1));
            if i == 10 {
                break;
            }
        }
    });

    handler.join().unwrap();

    println!("DONE");
}
