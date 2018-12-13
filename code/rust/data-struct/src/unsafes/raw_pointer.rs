pub fn demo_raw_pts() {
    let mut x = 10;
    let ptr_x = &mut x as *mut i32;

    let y = Box::new(20);
    let ptr_y = &*y as *const i32;

    unsafe {
        *ptr_x += *ptr_y;
    }

    println!("after {} x = {}", stringify!(*ptr_x += *ptr_y;), x);

    let null_ptr = ::std::ptr::null::<i32>();
    println!("{:?}", null_ptr);
}