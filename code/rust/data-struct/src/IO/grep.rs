use std::io::{self, BufRead, BufReader};
use std::fs::File;

fn grep<R>(target: &str, reader: R) -> io::Result<()> 
    where R: BufRead
{
    for line_result in reader.lines() {
        let line = line_result?;
        if line.contains(target) {
            println!("[GREP]:\t\t{}", line);
        }
    };
    Ok(())
}

pub fn demo_grep() {
    let target = "localhost";
    let f = File::open("/etc/hosts").unwrap();

    grep(&target, BufReader::new(f)).unwrap();
}