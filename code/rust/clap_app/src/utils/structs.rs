use std::fmt::{self, Display, Formatter};

#[derive(Debug)]
struct Point(f32, f32);

impl Display for Point {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        write!(f, "({}, {})", self.0, self.1)
    }
}

#[derive(Debug)]
struct Queue<T> {
    older: Vec<T>,
    younger: Vec<T>,
}

impl<T> Queue<T> {
    pub fn new() -> Queue<T> {
        Queue { older: Vec::new(), younger: Vec::new() }
    }
    pub fn push(&mut self, t: T) {
        self.younger.push(t);
    }
    pub fn is_empty(&self) -> bool {
        self.older.is_empty() && self.younger.is_empty()
    }
}

pub fn demo_structs() {
    let p = Point(2.0, 3.3);

    println!("{}", p);
    println!("{:?}", p);

    let mut q = Queue::new();
    let mut r = Queue::new();
    q.push("CAD");
    r.push(0.74);
    q.push("BTC");
    r.push(2737.7);
}