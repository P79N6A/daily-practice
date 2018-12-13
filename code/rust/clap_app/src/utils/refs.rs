use std::cell::{Cell,RefCell};
use std::fs::File;

struct SpiderBot {
    hardware_error_count: Cell<u32>,
    log_file: RefCell<File>,
}

impl SpiderBot {
    fn new() -> SpiderBot {
        let f = File::open("/etc/hosts").unwrap();
        SpiderBot {
            hardware_error_count: Cell::new(0),
            log_file: RefCell::new(f),
        }
    }

    fn add_hardware_error(&self) {
        let n = self.hardware_error_count.get();
        self.hardware_error_count.set(n+1);
    }
    fn hardware_errors(&self) -> u32 {
        self.hardware_error_count.get()
    }
    // fn log(&self, message: &str) {
    //     let mut file = self.log_file.borrow_mut();
    //     writeln!(file, "{}", message);
    // }
}

pub fn demo_refs() {
    let sb = SpiderBot::new();
    sb.add_hardware_error();

    println!("{}", sb.hardware_errors());
}