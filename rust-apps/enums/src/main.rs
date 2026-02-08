// #[allow(dead_code)]
// enum AboHava {
//     Aftabi,
//     Barani,
//     Barfi,
//     Abri,
// }

fn main() {
    // let today = AboHava::Abri;
    // match today {
    //     AboHava::Abri => {
    //         println!("AboHava is abri");
    //         println!("Naro biron ke hava del gire :(");
    //     }
    //     AboHava::Aftabi => {
    //         println!("AboHava is aftabi");
    //         println!("Boro biron ke hava alliye :)");
    //     }
    //     AboHava::Barani => {
    //         println!("AboHava is barani");
    //         println!("Hava doo nafaras heyf ke to single hasti :(");
    //     }
    //     AboHava::Barfi => {
    //         println!("AboHava is barfi");
    //         println!("Key berim barf bazi? :)");
    //     }
    //     _ => println!("No"),
    // }

    // let result = match div(50, 2) {
    //     Ok(v) => v,
    //     Err(e) => {
    //         eprintln!("error: {}", e);
    //         return;
    //     }
    // };

    let result = div(50, 20).expect("ERROR");
    println!("{}", result);
}

fn div(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        Err("No div to ziro".to_string())
    } else {
        Ok(a / b)
    }
}
