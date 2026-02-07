// struct Book {
//     title: String,
//     author: String,
// }

// impl Book {
//     fn print_info(self) {
//         println!("TitleBook: {}, Author: {}", self.title, self.author);
//     }
// }

fn main() {
    println!("Hi, my name is Pouya im programmer i love Golang, Rust, Python and Typescript.");

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
}
