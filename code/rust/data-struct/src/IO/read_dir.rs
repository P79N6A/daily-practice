use std::path::Path;
use std::fs::read_dir;

pub fn demo_read_dir() {
    let target = "/usr/";
    let path = Path::new(target);
    
    print!("{}:\t", target);

    for entry_list in path.read_dir().unwrap() {
        let entry = entry_list.unwrap();
        print!("{}\t", entry.file_name().to_string_lossy());
    }
    print!("\n");
}