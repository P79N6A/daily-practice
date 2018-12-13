use std::fmt::Debug;
use std::char::from_u32;

fn dump<T, U>(t: T)
where
    T: IntoIterator<Item = U>,
    U: Debug,
{
    for u in t {
        print!("{:?}", u);
    }
    println!("");
}

#[derive(Debug)]
struct MyCollection<T>(Vec<T>);

impl<T> MyCollection<T> {
    fn new() -> MyCollection<T> {
        MyCollection(Vec::<T>::new())
    }

    fn add(&mut self, elem: T) {
        self.0.push(elem);
    }
}

impl<T> IntoIterator for MyCollection<T> {
    type Item = T;
    type IntoIter = ::std::vec::IntoIter<T>;

    fn into_iter(self) -> Self::IntoIter {
        self.0.into_iter()
    }
}

pub fn demo_dump() {
    let mut c = MyCollection::new();

    for i in ('a' as u32)..('z' as u32) + 1 {
        c.add(from_u32(i).unwrap());
    }

    dump(c);
}
