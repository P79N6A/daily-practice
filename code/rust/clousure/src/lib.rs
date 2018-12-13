pub mod libs;

use std::thread;
use std::time::Duration;
use libs::cacher::Cacher;
use libs::counter::Counter;

fn generate_workout(intensity: u32, random_number: u32) {
    let mut expensive_result = Cacher::new(|num| {
        println!("calculating slowly...");
        thread::sleep(Duration::from_secs(2));
        num
    });

    if intensity < 25 {
       println!(
           "Today, do {} pushups!",
           expensive_result.value(intensity)
       );
       println!(
           "Next, do {} situps!",
           expensive_result.value(intensity)
       );
   } else {
       if random_number == 3 {
           println!("Take a break today! Remember to stay hydrated!");
       } else {
           println!(
               "Today, run for {} minutes!",
               expensive_result.value(intensity)
           );
       }
   }
}

impl Iterator for Counter {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        self.count += 1;

        if self.count < 6 {
            Some(self.count)
        } else {
            None
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn call_with_different_values() {
        let mut c = Cacher::new(|a| a);

        let v1 = c.value(1);
        let v2 = c.value(2);

        assert_eq!(v2, v1);
        assert_ne!(v2, 2);
    }

    #[test]
    fn calling_next_directly() {
        let mut counter = Counter::new();

        assert_eq!(counter.next(), Some(1));
        assert_eq!(counter.next(), Some(2));
        assert_eq!(counter.next(), Some(3));
        assert_eq!(counter.next(), Some(4));
        assert_eq!(counter.next(), Some(5));
        assert_eq!(counter.next(), None);
    }

    #[test]
    fn using_other_iterator_methods() {
        let sum: u32 = Counter::new().zip(Counter::new().skip(1))
                                     .map(|(a, b)| a * b)
                                     .filter(|x| x % 3 == 0)
                                     .sum();
        assert_eq!(18, sum);
    }
}
