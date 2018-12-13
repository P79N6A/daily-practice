use std::ops::Deref;
use std::rc::Rc;
use std::cell::RefCell;

pub mod mods;

#[derive(Debug)]
pub struct MyBox<T>(T);

impl<T> MyBox<T> {
    pub fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &T {
        &self.0
    }
}

#[derive(Debug)]
pub enum List {
    Cons(i32, Rc<List>),
    Nil,
}

use List::{Cons, Nil};

pub fn demo_strong_ref() {
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    println!("count after creating a = {}", Rc::strong_count(&a));
    let _b = Cons(3, Rc::clone(&a));
    println!("count after creating b = {}", Rc::strong_count(&a));
    {
        let _c = Cons(4, Rc::clone(&a));
        println!("count after creating c = {}", Rc::strong_count(&a));
    }
    println!("count after c goes out of scope = {}", Rc::strong_count(&a));
}

#[derive(Debug)]
pub enum List2 {
    Cons2(Rc<RefCell<i32>>, Rc<List2>),
    Nil2,
}

use List2::{Cons2, Nil2};

pub fn demo_ref_cell() {
    let value = Rc::new(RefCell::new(5));

    let a = Rc::new(Cons2(Rc::clone(&value), Rc::new(Nil2)));

    let b = Cons2(Rc::new(RefCell::new(6)), Rc::clone(&a));
    let c = Cons2(Rc::new(RefCell::new(10)), Rc::clone(&a));

    *value.borrow_mut() += 10;

    println!("a after = {:?}", a);
    println!("b after = {:?}", b);
    println!("c after = {:?}", c);
}


#[derive(Debug)]
pub enum List3 {
    Cons3(i32, RefCell<Rc<List3>>),
    Nil3,
}
impl List3 {
    fn tail(&self) -> Option<&RefCell<Rc<List3>>> {
        match *self {
            Cons3(_, ref item) => Some(item),
            Nil3 => None,
        }
    }
}
use List3::{Cons3, Nil3};

pub fn demo_rc_in_ref_cell() {
    let a = Rc::new(Cons3(5, RefCell::new(Rc::new(Nil3))));

    println!("a initial rc count = {}", Rc::strong_count(&a));
    println!("a next item = {:?}", a.tail());

    let b = Rc::new(Cons3(10, RefCell::new(Rc::clone(&a))));

    println!("a rc count after b creation = {}", Rc::strong_count(&a));
    println!("b initial rc count = {}", Rc::strong_count(&b));
    println!("b next item = {:?}", b.tail());

    if let Some(link) = a.tail() {
        *link.borrow_mut() = Rc::clone(&b);
    }

    println!("b rc count after changing a = {}", Rc::strong_count(&b));
    println!("a rc count after changing a = {}", Rc::strong_count(&a));

    // Uncomment the next line to see that we have a cycle; it will
    // overflow the stack
    // println!("a next item = {:?}", a.tail());
}
