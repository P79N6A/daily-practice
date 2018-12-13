use std::hash::Hash;
use std::borrow::Borrow;
use std::collections::HashMap;

struct Borrowable<K, V> {
    table: HashMap<K, V>,
}

impl<K, V> Borrowable<K, V>
where
    K: Eq + Hash,
{
    fn get<Q: ?Sized>(&self, key: &Q) -> Option<&V>
    where
        K: Borrow<Q>,
        Q: Eq + Hash,
    {
        self.table.get(key)
    }

    fn new() -> Borrowable<K, V> {
        Borrowable {
            table: HashMap::<K, V>::new(),
        }
    }

    fn insert<Q: ?Sized>(&mut self, key: &Q, value: V)
    where
        Q: ToOwned<Owned = K>,
        K: Borrow<Q>,
    {
        self.table.insert(key.to_owned(), value);
    }
}

pub fn demo_borrowable() {
    let mut b = Borrowable::<String, &str>::new();
    b.insert("name", "Jack");
    b.insert("age", "23");
    b.insert("sex", "male");

    println!("{:?}", b.get("name"));
}
