use std::collections::VecDeque;
use std::{thread, time};
use tokio::prelude::*;

const SLEEP_MILLIS: u64 = 100;
fn poll_number() -> Async<u32> {
    Async::Ready(233)
}
pub struct MyTask;
impl Future for MyTask {
    type Item = ();
    type Error = ();

    fn poll(&mut self) -> Result<Async<()>, ()> {
        match poll_number() {
            Async::Ready(number) => {
                println!("number={:?}", number);
                Ok(Async::Ready(()))
            }
            Async::NotReady => {
                return Ok(Async::NotReady);
            }
        }
    }
}

pub struct SpinExecutor {
    ready_tasks: VecDeque<Box<Future<Item = (), Error = ()>>>,
    not_ready_tasks: VecDeque<Box<Future<Item = (), Error = ()>>>,
}
impl SpinExecutor {
    pub fn spawn<T>(&mut self, task: T)
    where T: Future<Item = (), Error = ()> + 'static {
        self.not_ready_tasks.push_back(Box::new(task));
    }
    pub fn run(&mut self) {
        loop {
            while let Some(mut task) = self.ready_tasks.pop_front() {
                match task.poll().unwrap() {
                    Async::Ready(_) => {}
                    Async::NotReady => {
                        self.not_ready_tasks.push_back(task);
                    }
                }
            }
            if self.not_ready_tasks.is_empty() { return; }
            self.sleep_until_tasks_are_ready();
        }
    }
    fn sleep_until_tasks_are_ready(&mut self) {
        loop {
            if self.not_ready_tasks.is_empty() {
                thread::sleep(time::Duration::from_millis(SLEEP_MILLIS));
            } else {
                while let Some(mut task) = self.not_ready_tasks.pop_front() {
                    self.ready_tasks.push_back(task);    
                }
                return;
            }
        }
    }
}