extern crate rand;

use std::thread::{spawn, JoinHandle, sleep};
use std::sync::mpsc::{Sender, Receiver, channel, IntoIter};
use std::time::{SystemTime, UNIX_EPOCH, Duration};
use self::rand::Rng;
use self::rand::distributions::{IndependentSample, Range};

fn now_secs() -> u64 {
    SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs()
}

fn generate(threads: u64) -> IntoIter<u64> {
    let (sender, receiver): (Sender<u64>, Receiver<u64>) = channel();

    for i in 0..threads {
        let sender_local = sender.clone();
        spawn(move || {
            let mut rng = rand::thread_rng();
            let random_door = Range::new(1, 3);
            let secs = random_door.ind_sample(&mut rng); // sleep randomly
            sleep(Duration::from_secs(secs));
            sender_local.send(i + now_secs());
        });
    }

    receiver.into_iter()
}

pub fn demo_interaction() {
    for r in generate(12) {
        println!("received: {}", r);
    }
}