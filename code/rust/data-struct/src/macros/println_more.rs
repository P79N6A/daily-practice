#[macro_export]
macro_rules! println_more {
    ($fmt:expr) => {
        use std::io::{Write, stdout};
        println!("[println_more!]: {} 行 {} 列 {}", file!(), line!(), column!());
        stdout().write_fmt(format_args!($fmt));
    }; 
    ($fmt:expr, $($args:tt),*) => {
        use std::io::{Write, stdout};
        println!("[println_more!]: {} 行 {} 列 {}", file!(), line!(), column!());
        stdout().write_fmt(format_args!($fmt, $($args),*));
    };
}
