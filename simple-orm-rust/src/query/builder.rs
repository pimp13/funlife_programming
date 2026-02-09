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

impl QueryBuilder {
    pub fn new(table: &str) -> Self {
        QueryBuilder {
            table: table.to_string(),
            condition: Vec::new(),
            order_by: None,
            limit: None,
        }
    }

    pub fn where_eq(mut self, field: &str, value: &str) -> Self {
        self.condition.push(Condition {
            field: field.to_string(),
            operator: Operator::Eq,
            value: value.to_string(),
        });
        self
    }

    pub fn where_gt(mut self, field: &str, value: &str) -> Self {
        self.condition.push(Condition {
            field: field.to_string(),
            operator: Operator::Gt,
            value: value.to_string(),
        });
        self
    }

    pub fn order_by(mut self, field: &str, dir: &str) -> Self {
        self.order_by = Some((field.to_string(), dir.to_string()));
        self
    }

    pub fn limit(mut self, limit: u32) -> Self {
        self.limit = Some(limit);
        self
    }

    pub fn build(&self) -> (String, Vec<String>) {
        let mut sql = format!("SELECT * FROM {}", self.table);
        let mut params = Vec::new();

        if !self.condition.is_empty() {
            let parts: Vec<String> = self
                .condition
                .iter()
                .enumerate()
                .map(|(i, c)| {
                    params.push(c.value.clone());
                    c.to_sql(i + 1)
                })
                .collect();

            sql.push_str(" WHERE ");
            sql.push_str(&parts.join(" AND "));
        }

        if let Some((field, dir)) = &self.order_by {
            sql.push_str(&format!(" ORDER BY {} {}", field, dir));
        }

        if let Some(limit) = self.limit {
            sql.push_str(&format!(" LIMIT {}", limit));
        }

        (sql, params)
    }
}
