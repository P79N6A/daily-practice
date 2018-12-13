use std::str::FromStr;
use std::num::ParseIntError;

#[derive(Debug, PartialEq)]
struct Point {
    x: u32,
    y: u32,
}
impl FromStr for Point {
    type Err = ParseIntError;
    
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let coords: Vec<&str> = s.trim_matches(|c| c == '(' || c == ')')
            .split(",")
            .collect();
        
        let x = coords[0].trim().parse::<u32>()?;
        let y = coords[1].trim().parse::<u32>()?;
        
        Ok(Point{ x, y })
    }
}

pub fn demo_from_str() {  
    let p = Point::from_str("(11 , 23)");
    
    println!("{:?}", p);
}