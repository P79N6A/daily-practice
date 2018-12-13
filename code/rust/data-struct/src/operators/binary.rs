use std::ops::Add;

#[derive(Debug)]
struct Slices<T> {
    data: Vec<Box<T>>,
}

impl<T> Slices<T>
where
    T: Clone,
{
    fn from_vec(vec: &Vec<T>) -> Slices<T> {
        let mut data = Vec::new();
        vec.iter()
            .cloned()
            .for_each(|elem| data.push(Box::new(elem)));
        Slices::new(data)
    }

    fn new(data: Vec<Box<T>>) -> Slices<T> {
        Slices { data }
    }
}

impl<T> Add for Slices<T>
where
    T: Add<Output = T> + Copy,
{
    type Output = Slices<T>;

    fn add(self, other: Self::Output) -> Self::Output {
        let mut data = Vec::new();
        for i in 0..self.data.len() {
            let t = *self.data[i] + *other.data[i];
            data.push(Box::new(t));
        }
        Slices::new(data)
    }
}

pub fn demo_binary() {
    let v1 = vec![1; 3];
    let v2 = vec![2; 3];
    let a = Slices::from_vec(&v1);
    let b = Slices::from_vec(&v2);

    println!("{:?}", a + b);
}
