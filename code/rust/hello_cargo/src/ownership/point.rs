use std::fmt::Debug;

#[derive(Debug, PartialEq)]
pub struct Point<T, U> {
    pub x: T,
    pub y: U,
}
impl<T, U> Point<T, U>{
    pub fn x(&self) -> &T {
        &self.x
    }
    pub fn y(&self) -> &U {
        &self.y
    }

    pub fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}

pub trait Summarizable {
    fn summary(&self) -> String;

    fn greeting(&self) -> String where Self: Debug {
        format!("Ref to: [{}]", self.summary())
    }
}
impl<T: Debug, U: Debug> Summarizable for Point<T, U> {
    fn summary(&self) -> String {
        format!("Point: {} x:{:?}, y:{:?} {}", "{", self.x(), self.y(), "}")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_summary() {
        let s1 = Point{ x: 4, y: 5.2 };
        let s2 = s1.summary();
        let s3 = format!("Point: {} x:{:?}, y:{:?} {}", "{", s1.x(), s1.y(), "}");
        let s4 = String::from("casual string");
        assert_eq!(s2, s3);
        assert_ne!(s2, s4);
    }
}
