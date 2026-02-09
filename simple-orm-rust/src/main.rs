mod query;

use crate::query::builder::QueryBuilder;

fn main() {
    let (sql, params) = QueryBuilder::new("users")
        .where_eq("status", "active")
        .where_gt("age", "16")
        .order_by("created_at", "DESC")
        .limit(10)
        .build();

    println!("SQL: {} , Params: {:#?}", sql, params);
}
