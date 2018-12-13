use std::io::{self, BufRead, BufReader};
use std::fs::File;

pub fn demo_collect_lines() {
    let target = "/etc/hosts";
    let f = File::open(target).unwrap();
    let reader = BufReader::new(f);

    let lines = reader.lines().collect::<io::Result<Vec<String>>>();
    println!("{}: {:?}", target, lines.unwrap());
}