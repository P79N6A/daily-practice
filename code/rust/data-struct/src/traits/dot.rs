use std::ops::{Add, Mul};

fn dot<N>(v1: &[N], v2: &[N]) -> N
where
    N: Add<Output = N> + Mul<Output = N> + Default + Copy,
{
    let mut total = N::default();
    for i in 0..v1.len() {
        total = total + v1[i] * v2[i];
    }
    total
}

pub fn demo_dots() {
    let v1 = [1, 2, 3, 4];
    let v2 = [1, 1, 1, 1];
    println!("dot(&{:?}, &{:?}) = {}", &v1, &v2, dot(&v1, &v2));
}
