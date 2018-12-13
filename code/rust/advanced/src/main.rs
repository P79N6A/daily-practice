extern crate advanced;

use advanced::mods::supertraits::{Point, OutlinePrint};
use advanced::mods::wrapper::Wrapper;
use advanced::mods::newtypes;

fn main() {
    let p = Point{ x:2, y:3 };
    p.outline_print();

    let w = Wrapper(vec![String::from("hello"), String::from("world")]);
    println!("w = {}", w);

    newtypes::demo_newtypes();
}
