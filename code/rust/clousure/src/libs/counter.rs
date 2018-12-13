pub struct Counter {
    pub count: u32,
}

impl Counter {
    pub fn new() -> Counter {
        Counter { count: 0 }
    }
}
