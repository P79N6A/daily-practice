use std::fmt::Debug;

pub fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

pub fn longest2<'a, T>(x: &'a str, y: &'a str, ann: T) -> &'a str
    where T: Debug
{
    println!("Announcement! {:?}", ann);
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_longest() {
        let x = "hello";
        let y = "haha";
        let z = longest(x, y);
        assert_eq!("hello", z);
    }
}
