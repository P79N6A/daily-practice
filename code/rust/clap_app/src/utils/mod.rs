pub mod collection;
pub mod publication;
pub mod fault;
pub mod structs;
pub mod refs;

use std::collections::HashMap;
use std::fmt::Display;

pub type Table = HashMap<String, Vec<String>>;

pub fn show(table: &Table) {
    for (artist, works) in table {
        println!("works by {}", artist);

        for work in works {
            println!("  {}", work)
        }
    }
}

pub fn print<T>(c: T) where T: Display{
    println!("{}", c)
}

pub fn demo_shadowing() {
    let x = 6;
    {
        let x = 12;
        print(x);
    }
    print(x);
}

pub struct StringTable {
    pub elements: Vec<String>,
}

impl StringTable {
    pub fn new(elements: Vec<&str>) -> StringTable {
        let elements = collection::from_strs(elements);
        StringTable {
            elements,
        }
    }
     
    pub fn find_by_prefix<'a, 'b>(&'a self, prefix: &'b str) -> Option<&'a String> {
        let ps = String::from(prefix).to_lowercase();
        let pss = ps.as_str();
        for i in 0..self.elements.len() {
            if self.elements[i].to_lowercase().starts_with(pss) {
                return Some(&self.elements[i])
            }
        }
        None
    }

    pub fn demo_finding() {
        let st = StringTable::new(vec!("Yes", "No", "Haha"));
        match st.find_by_prefix("ha") {
            Some(word) => {
                println!("{}", word)
            }
            None => println!("Not found")
        }
    }
}