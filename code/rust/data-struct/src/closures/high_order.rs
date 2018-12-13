fn pow(x: i32) -> impl Fn(u32) -> i32 {
    move |y| x.pow(y)
}

pub fn demo_high_order() {
    let x = 2;
    let y = 3;
    println!("{} ** {} = {}", x, y, pow(x)(y));
}