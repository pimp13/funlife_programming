#[derive(Debug)]
struct Order {
    id: u32,
    amount: f64,
    status: String,
}

fn main() {
    // let numbers = vec![1, 2, 3, 4, 5];

    // let mut iter = numbers.iter();
    // println!("{:?}", iter.next());
    // println!("{:?}", iter.next());
    // println!("{:?}", iter.next());
    // println!("{:?}", iter.next());
    // println!("{:?}", iter.next());
    // println!("{:?}", iter.next());

    // let orders = vec![
    //     Order {
    //         id: 1,
    //         amount: 120.0,
    //         status: String::from("Complited"),
    //     },
    //     Order {
    //         id: 2,
    //         amount: 320.0,
    //         status: String::from("Pending"),
    //     },
    //     Order {
    //         id: 3,
    //         amount: 240.0,
    //         status: String::from("Complited"),
    //     },
    // ];

    // let complited_order: Vec<&Order> = orders
    //     .iter()
    //     .filter(|order| order.status == "Complited")
    //     .collect();

    // for order in complited_order {
    //     println!(
    //         "OrderID: {} , Amount: ${} , Status: {}",
    //         order.id, order.amount, order.status
    //     )
    // }
}
