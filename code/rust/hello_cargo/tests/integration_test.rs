extern crate hello_cargo;
use hello_cargo::ownership;

mod common;

#[test]
fn should_get_longest_word() {
    common::setup();

    let x = "good";
    let y = "afternoon";
    let z = ownership::strings::longest(x, y);
    assert_eq!(y, z);
}

#[test]
fn should_raw_type_eq() {
    assert_eq!(b'A', 65u8);

    assert_eq!( 10_i8 as u16, 10_u16); // in range
    assert_eq!( 2525_u16 as i16, 2525_i16); // in range
    assert_eq!( -1_i16 as i32, -1_i32); // sign-extended
    assert_eq!(65535_u16 as i32, 65535_i32); // zero-extended
    // Conversions that are out of range for the destination
    // produce values that are equivalent to the original modulo 2^N,
    // where N is the width of the destination in bits. This
    // is sometimes called "truncation".
    assert_eq!( 1000_i16 as u8, 232_u8);
    assert_eq!(65535_u32 as i16, -1_i16);
    assert_eq!( -1_i8 as u8, 255_u8);
    assert_eq!( 255_u8 as i8, -1_i8);

    assert!((-1. / std::f32::INFINITY).is_sign_negative());
}

struct Point {
    x: i32,
    y: i32,
}

#[test]
fn destruct_struct() {
    let p = Point { x: 0, y: 7 };

    let Point { x, y } = p;
    assert_eq!(0, x);
    assert_eq!(7, y);
}
