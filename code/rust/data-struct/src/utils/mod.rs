use std::fmt::Debug;
use std::fmt::Binary;

pub fn println<T>(a: T)
where
    T: Debug,
{
    print!("{:?}\n", a);
}

pub fn println_binary<T>(a: T)
where
    T: Binary,
{
    print!("{:#b}\n", a);
}
