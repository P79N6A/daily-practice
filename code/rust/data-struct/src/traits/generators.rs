use std::ops::{Generator, GeneratorState};

pub fn demo_generator() {
    let mut generator = || {
        let mut curr : u64 = 1;
        let mut next : u64 = 1;
        loop {
            let new_next = curr.checked_add(next);
            if let Some(new_next) = new_next {
                curr = next;
                next = new_next;
                yield curr;
            } else {
                return;
            }
        }
    };
    loop {
        unsafe {
            match generator.resume() {
                GeneratorState::Yielded(v) => println!("{}", v),
                GeneratorState::Complete(_) => return,
            }
        }
    }
}