use std::sync::{Arc, Mutex, Condvar};
use std::thread::spawn;

pub fn demo_condvar() {
    let pair = Arc::new((Mutex::new(false), Condvar::new()));
    let pair2 = pair.clone();

    spawn(move || {
        let (ref lock, ref cvar) = *pair2;
        let mut started = lock.lock().unwrap();
        *started = true;
        println!("[CONDVAR] notify_one");
        cvar.notify_one();
    });

    let (ref lock, ref cvar) = *pair;
    let mut started = lock.lock().unwrap();
    while !*started {
        println!("[CONDVAR] !started");
        started = cvar.wait(started).unwrap();
    }
    println!("[CONDVAR] OVER!");
}

