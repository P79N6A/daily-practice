type Thunk = Box<Fn() + Send + 'static>;

pub fn demo_newtypes() {
    let f: Thunk = Box::new(|| println!("hello"));

    f();
}
