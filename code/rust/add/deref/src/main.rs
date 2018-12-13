extern crate deref;

use deref::MyBox;
use deref::mods::node;

fn main() {
    let x = 5;
    let y = MyBox::new(x);

    println!("{}", x);
    println!("{:?}", y);
    println!("{}", *y);

    deref::demo_strong_ref();

    deref::demo_ref_cell();

    deref::demo_rc_in_ref_cell();

    node::demo_tree();

    node::demo_visual_count();
}
