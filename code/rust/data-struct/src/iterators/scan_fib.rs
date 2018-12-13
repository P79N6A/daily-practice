fn fib_vec(pos: i32) -> Option<(i32, i32)> {
    match pos {
        _ if pos < 1 => None,
        _ => {
            let iter = (0..pos).scan((1, 1), |sum, _curr| {
                let s = sum.0 + sum.1;
                sum.0 = sum.1;
                sum.1 = s;
                Some(*sum)
            });
            iter.last()
        }
    }
}

pub fn demo_scan_fib_vec() {
    let pos = 10;
    let pair = fib_vec(pos);

    println!(
        "Fibonacci({}) = {:?}",
        pos,
        match pair {
            Some(tuple) => tuple.0,
            None => 0,
        }
    );
}
