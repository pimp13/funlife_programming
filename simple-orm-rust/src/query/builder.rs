#[derive(Debug, Clone)]
pub enum Operator {
    Eq,
    Gt,
    Lt,
}

impl Operator {
    fn as_sql(&self) -> &'static str {
        match self {
            Operator::Eq => "=",
            Operator::Gt => ">",
            Operator::Lt => "<",
        }
    }
}

#[derive(Debug, Clone)]
pub struct Condition {
    pub field: String,
    pub operator: Operator,
    pub value: String,
}

impl Condition {
    pub fn to_sql(&self, index: usize) -> String {
        format!("{} {} ${}", self.field, self.operator.as_sql(), index)
    }
}

#[derive(Debug)]
pub struct QueryBuilder {
    table: String,
    condition: Vec<Condition>,
    order_by: Option<(String, String)>,
    limit: Option<u32>,
}
