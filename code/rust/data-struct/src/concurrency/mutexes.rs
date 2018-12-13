use std::sync::{Arc, Mutex};
use std::thread;
use std::sync::mpsc::channel;

const N: usize = 10;

pub fn demo_mutexes() {
    let data = Arc::new(Mutex::new(666));
    let lock = data.clone();
    let (tx, rx) = channel();    

    for _ in 0..N {
        let (data, tx) = (data.clone(), tx.clone());
        thread::spawn(move || {
            let mut data = match data.lock() {
                Ok(data) => data,
                Err(poisoned) => poisoned.into_inner(),
            };
            *data += 1;
            if *data == N {
                tx.send(()).unwrap();
            }
            panic!();
        }).join();
    }

    for _ in 0..N {
        match lock.lock() {
            Ok(t) => println!("rx:{:?}", t),
            Err(poisoned) => println!("posioned:{:?}", *(poisoned.into_inner())),
        }
    }
}

