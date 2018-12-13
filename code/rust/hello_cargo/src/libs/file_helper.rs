use std::io;
use std::io::Read;
use std::fs::File;

pub fn read_file(filename: &str) -> Result<String, io::Error> {
    let mut f = File::open(filename)?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}
