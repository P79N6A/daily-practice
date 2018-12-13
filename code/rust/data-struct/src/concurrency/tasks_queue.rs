extern crate rand;

use std::thread::{self, spawn, JoinHandle};
use std::collections::VecDeque;
use std::io::Result;
use std::time::Duration;
use self::rand::Rng;

trait FnBox{
    fn call_box(self: Box<Self>);
}
impl<F: FnOnce()> FnBox for F {
    fn call_box(self: Box<Self>) {
        (*self)();
    }
}
type Job = Box<FnBox + Send + 'static>;
enum Envelope {
    NewJob(Job),
}

struct Tasks {
    queue: Vec<Envelope>,
    tasks: Vec<JoinHandle<()>>,
}
impl Tasks {
    fn new() -> Tasks { 
        Tasks { queue: Vec::new(), tasks: Vec::new() } 
    }
    fn add(&mut self, f: Envelope) {
        self.queue.push(f);      
    }
    fn run(mut self) -> Result<()> {
        for f in self.queue.into_iter() {
            match f {
                Envelope::NewJob(job) => {
                    self.tasks.push(spawn(move || {
                        job.call_box()
                    }));
                }
            } 
        }
        for task in self.tasks.into_iter() {
            task.join().unwrap();
        }
        Ok(())
    }
}

fn make_task() -> Envelope {
    let job = Box::new(move || {
        let mut rng = rand::thread_rng();
        let start = rng.gen::<u32>();
        let thread_id = thread::current().id();
        println!("[{:<2?}]: got number: {}", thread_id, start);
    });
    Envelope::NewJob(job)
}

pub fn demo_tasks() {
    let mut queue = Tasks::new();

    for i in 0..10 {
        queue.add(make_task());   
    }
    
    queue.run();
}