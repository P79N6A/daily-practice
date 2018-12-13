trait Events {
    type Item;

    fn born(&self) -> Option<Self::Item>;
    fn die(&self) -> Option<Self::Item>;
}

struct Person<'a> {
    name: &'a str,
}

impl<'a> Events for Person<'a> {
    type Item = String;

    fn born(&self) -> Option<Self::Item> {
        Some(format!("{} was Born", self.name))
    }
    fn die(&self) -> Option<Self::Item> {
        Some(format!("{} was Died", self.name))
    }
}

pub fn demo_associated_type() {
    let peter = Person { name: "Peter" };
    println!("EVENT: {}", peter.born().unwrap());
    println!("EVENT: {}", peter.die().unwrap());
}
