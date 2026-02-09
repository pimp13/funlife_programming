// struct Book {
//     title: String,
//     author: String,
// }

// impl Book {
//     fn print_info(self) {
//         println!("TitleBook: {}, Author: {}", self.title, self.author);
//     }
// }

use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut counts: HashMap<i32, i32> = HashMap::new();

        for n in nums {
            let or_insert = counts
                .entry(n)
                .and_modify(|counter| *counter += 1)
                .or_insert(1);
        }

        42
    }
}

fn main() {
    // println!("Hi, my name is Pouya im programmer i love Golang, Rust, Python and Typescript.");
    // Lifetime in rust
    /*
    error mideh: borrowed value does not live long enough
    dar code payien variable z tarif shode va dar scoop variable x tarif shode
    va refrence x ro yani meghdar x ro daram mirizam dar z
    vali error mideh chon ke meghdar x ziyad zendeh nemimofe va chon dar scoop hast sari mimire

    life time ba <'a> tarif mishe
    */
    // let z;
    // {
    //     let x = 42;
    //     z = &x;
    // }

    // let my_book = Book {
    //     title: String::from("First My Book"),
    //     author: String::from("DevPouyaGh"),
    // };
    // my_book.print_info();

    let nums = vec![1, 2, 2, 3, 3, 3, 4, 5];
    let result = Solution::majority_element(nums);
    println!("{}", result);
}
