use std::ops::{Add, Sub};

#[derive(Debug, PartialEq)]
struct Point {
    x: i32,
    y: i32,
}

impl Add for Point {
    type Output = Point;

    fn add(self, other: Point) -> Point {
        Point {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}
impl Sub for Point {
    type Output = Point;

    fn sub(self, other: Point) -> Point {
        Point {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

struct Millimeters(u32);
struct Meters(u32);

impl Add<Meters> for Millimeters {
    type Output = Millimeters;

    fn add(self, other: Meters) -> Millimeters {
        Millimeters(self.0 + (other.0 * 1000))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn operator_add_should_eq() {
        assert_eq!(Point{ x: 1, y: 0 } + Point{ x: 2, y: 3 },
                   Point{ x: 3, y: 3 });
    }

    #[test]
    fn operator_sub_should_eq() {
        assert_eq!(Point { x: 3, y: 3 } - Point { x: 2, y: 3 },
                   Point { x: 1, y: 0 });
    }
}
