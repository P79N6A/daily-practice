macro_rules! double_method {
    ($self_:ident, $body:expr) => {
        fn double(mut $self_) -> Dummy {
            $body
        }
    };
}

#[derive(Debug)]
struct Dummy(i32);
impl Dummy {
    double_method! {self, {
        self.0 *= 2;
        self
    }}
}

pub fn demo_ident_self() {
    let d = Dummy(16);
    println!("{:?}", d.double());
}