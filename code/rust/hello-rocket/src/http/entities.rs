#[derive(Debug, Serialize)]
pub struct Person<'a> {
    pub name: &'a str,
    pub sex: &'a str,
    pub age: u16,
}
