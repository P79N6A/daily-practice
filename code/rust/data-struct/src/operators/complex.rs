use std::ops::Add;

#[derive(Clone, Copy, Debug)]
struct Complex<T> {
    re: T,
    im: T,
}

impl<L, R, O> Add<Complex<R>> for Complex<L>
where
    L: Add<R, Output = O>,
{
    type Output = Complex<O>;
    fn add(self, rhs: Complex<R>) -> Self::Output {
        Complex {
            re: self.re + rhs.re,
            im: self.im + rhs.im,
        }
    }
}

pub fn demo_complexes() {
    let a = Complex::<i32> { re: 333, im: 444 };
    let b = Complex::<i32> { re: 111, im: 222 };

    println!("{:?}", a + b);
}
