extern crate holy_crate;

use holy_crate::kinds::PrimaryColor;
use holy_crate::utils::mix;

fn main() {
    let red = PrimaryColor::Red;
    let yellow = PrimaryColor::Yellow;
    mix(red, yellow);
}
