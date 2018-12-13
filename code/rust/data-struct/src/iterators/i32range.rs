use utils::*;

struct I32Range {
    start: i32,
    end: i32,
}

impl I32Range {
    fn new(start: i32, end: i32) -> Self {
        I32Range { start, end }
    }
}

impl Iterator for I32Range {
    type Item = i32;
    fn next(&mut self) -> Option<i32> {
        if self.start >= self.end {
            return None;
        }
        let result = Some(self.start);
        self.start += 1;
        result
    }
}

pub fn demo_i32range() {
    let r = I32Range::new(3, 22);
    for i in r {
        print!("{} ", i);
    }
    println("");
}
