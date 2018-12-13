extern crate rand;

pub mod lock;
pub mod oo;
pub mod unsafety;

use rand::Rng;
use std::thread;
use std::time::Duration;
use std::sync::mpsc;

pub fn demo_spawn_threads() {
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1));
        }

        "hello"
    });

    for i in 1..5 {
        println!("hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1));
    }

    println!("result from spawned thread: {}", handle.join().unwrap());
}

pub fn demo_threads_with_closure() {
    let v = vec![1, 2, 3];

    let handle = thread::spawn(move || {
        println!("Here's a vector: {:?}", v);
    });

    handle.join().unwrap();
}

pub fn demo_channel() {
    let (tx, rx) = mpsc::channel();
    let tx1 = mpsc::Sender::clone(&tx);

    thread::spawn(move || {
        let vals = vec![
            String::from("厉害了"),
            String::from("我的"),
            String::from("锅"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_millis(rand::thread_rng().gen_range(100, 999)));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("what"),
            String::from("the"),
            String::from("fuck"),
        ];

        for val in vals {
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_millis(rand::thread_rng().gen_range(100, 999)));
        }
    });

    for received in rx {
        println!("recv: {}", received);
    }
}
