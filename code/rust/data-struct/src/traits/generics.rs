use std::fmt::{self, Debug, Display, Formatter};

trait Crap {
    fn crap(&self);
}
// extend trait: Debug
impl<D: Debug> Crap for D {
    fn crap(&self) {
        println!("Quack Quack: {:?}", *self);
    }
}

// qualify target by trait
fn log<T>(obj: T)
where
    T: Debug + Display,
{
    println!("{:?}", obj);
}

// demostration
#[derive(Debug, Clone)]
struct Holy {
    text: String,
}
impl Display for Holy {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        write!(f, "Holy({})", self.text)
    }
}

pub fn demo_generics() {
    let h = Holy {
        text: "holy crap".to_string(),
    };
    log(h.clone());
    h.crap();
}
