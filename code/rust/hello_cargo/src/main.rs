pub mod ownership;
pub mod libs;

use ownership::point;
use ownership::point::Summarizable;
use ownership::strings;

fn main() {
    let s1 = point::Point{ x: 4, y: 5.2 };
    let s2 = point::Point{ x: 1.2, y: "hello" };

    {
        let s4 = &s2;
        println!("{:?}", s4.greeting());
    }

    let s3 = s1.mixup(s2);
    println!("{:?}", s3.summary());

    let string1 = String::from("abcd");
    let result;
    {
        let string2 = "xyz";
        result = strings::longest2(string1.as_str(), string2, ['1', '2', '3']);
    }
    println!("The longest string is {}", result);
}
