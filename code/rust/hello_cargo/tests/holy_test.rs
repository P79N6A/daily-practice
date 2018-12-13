extern crate hello_cargo;
use hello_cargo::ownership;

#[test]
fn should_get_longest_word() {
    let x = "good";
    let y = "afternoon";
    let z = ownership::strings::longest(x, y);
    assert_eq!(y, z);
}