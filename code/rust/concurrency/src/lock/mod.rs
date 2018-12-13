use std::thread;
use std::sync::{Mutex, Arc};

pub fn demo_mutex() {
    let m = Mutex::new(5);

    {//RAII lock
        let mut num = m.lock().unwrap();
        *num = 6;
    }

    println!("m = {:?}", m);
}

pub fn demo_threads_join() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();

            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Results: {}", *counter.lock().unwrap());
}
