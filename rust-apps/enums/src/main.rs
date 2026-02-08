#[allow(dead_code)]
enum AboHava {
    Aftabi,
    Barani,
    Barfi,
    Abri,
}

fn main() {
    let today = AboHava::Abri;
    match today {
        AboHava::Abri => {
            println!("AboHava is abri");
            println!("Naro biron ke hava del gire :(");
        }
        AboHava::Aftabi => {
            println!("AboHava is aftabi");
            println!("Boro biron ke hava alliye :)");
        }
        AboHava::Barani => {
            println!("AboHava is barani");
            println!("Hava doo nafaras heyf ke to single hasti :(");
        }
        AboHava::Barfi => {
            println!("AboHava is barfi");
            println!("Key berim barf bazi? :)");
        } // Default
          // _ => println!("No"),
    }
}
