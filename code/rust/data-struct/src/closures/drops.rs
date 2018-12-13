#[derive(Debug)]
struct Holy {
    crap: String,
}
impl Holy {
    fn new(s: &str) -> Holy {
        Holy { crap: s.to_owned() }
    }
}
impl Drop for Holy {
    fn drop(&mut self) {
        println!("{:?} is dying", self);
    }
}

pub fn demo_drops() {
    let h = Holy::new("Aloha");
    drop(h);

    println!("{}", "Emmmm....");
}
