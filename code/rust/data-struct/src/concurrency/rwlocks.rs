use std::sync::{Arc, RwLock};
use std::collections::HashMap;
use std::thread::{spawn, JoinHandle};

fn init_status() -> RwLock<HashMap<String, String>> {
    let mut h = HashMap::new();
    h.insert("name".to_string(), "Mark".to_string());
    h.insert("sex".to_string(), "Male".to_string());

    RwLock::new(h)
}

fn read(l: Arc<RwLock<HashMap<String, String>>>) -> JoinHandle<()> {
    spawn(move || {
        let h = l.read().unwrap();
        println!("[READ ]: name: {}", (*h).get("name").unwrap());
    })    
}

fn write(l: Arc<RwLock<HashMap<String, String>>>) -> JoinHandle<()>{
    spawn(move || {
        let mut h = l.write().unwrap();
        (*h).insert("name".to_string(), "Holy".to_string());
        println!("[WRITE]: name: Holy");
    })    
}

pub fn demo_rwlocks() {
    let status = Arc::new(init_status());
    let mut handlers = Vec::new();

    for i in 0..6 {
        let s = status.clone();
        match i {
            2 => handlers.push(write(s)),
            _ => handlers.push(read(s)),
        }
    }

    for handler in handlers {
        handler.join();
    }
}