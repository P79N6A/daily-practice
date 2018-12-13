extern crate rand;

use std::sync::Arc;
use std::sync::atomic::{AtomicIsize, Ordering};
use std::thread::{spawn, JoinHandle};
use self::rand::Rng;

fn write(atomic: Arc<AtomicIsize>) -> JoinHandle<()>{
    spawn(move || {
        let mut rng = rand::thread_rng();
        let i = rng.gen::<isize>();
        atomic.store(i, Ordering::SeqCst);
        println!("[ATOMIC] WRITE i = {}", i);
    })
}

fn read(atomic: Arc<AtomicIsize>) -> JoinHandle<()>{
    spawn(move || {
        println!("[ATOMIC] READ i = {}", atomic.load(Ordering::SeqCst));
    })
}

pub fn demo_atomics() {
    const N: usize = 6;
    let atomic = Arc::new(AtomicIsize::new(666));
    let mut handlers = Vec::new();

    for i in 0..N {
        let atomic_clone = atomic.clone();
        match i {
            i if i % 2 == 0 => handlers.push(write(atomic_clone)),
            _ => handlers.push(read(atomic_clone)),
        }
    }

    for handler in handlers {
        handler.join();
    }
}
