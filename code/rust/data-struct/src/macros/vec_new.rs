#[macro_export]
macro_rules! vec_new {
    ($elem:expr ; $n:expr) => {
        ::std::vec::Vec::new($elem, $n)
    };
    ($( $x:expr ),*) => {
        <[_]>::into_vec(Box::new([ $($x),* ]))
    };
    ($( $x:expr ),+ ,) => {
        vec_new![ $( $x ),* ]
    };
}
