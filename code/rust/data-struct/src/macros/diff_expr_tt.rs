macro_rules! stringify_expr {
    ($e:expr) => {
        stringify!($e)
    };
}

macro_rules! stringify_token_tree {
    ($($t:tt)*) => {
        stringify!($($t)*)
    };
}

pub fn demo_difference_of_expr_and_tt() {
    println!("{:?}", stringify_expr!(dummy(2 * (1 + (3)))));
    println!("{:?}", stringify_token_tree!(dummy(2 * (1 + (3)))));
}