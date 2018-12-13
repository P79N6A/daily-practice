use std::slice;

fn split_at_mut(slice: &mut [i32], mid: usize) -> (&mut [i32], &mut [i32]) {
    let len = slice.len();
    let ptr = slice.as_mut_ptr();

    assert!(mid <= len);

    unsafe {
        (slice::from_raw_parts_mut(ptr, mid),
         slice::from_raw_parts_mut(ptr.offset(mid as isize), len - mid))
    }
}

pub fn demo_unsafe_raw_pointer_of_slice() {
    let mut v = vec![1, 2, 3, 4, 5, 6];
    let r = &mut v[..];

    let (a, b) = split_at_mut(r, 3);
    println!("{:?}, {:?}", a, b);
}

extern "C" {
    fn abs(input: i32) -> i32;
}

pub fn demo_ffi() {
    unsafe {
        println!("Absolute value of -3 according to C: {}", abs(-3));
    }
}
