use std::intrinsics;

fn get_type<T>(_: T) -> String {
    unsafe {
        return String::from(intrinsics::type_name::<T>());
    }
}

fn print_type_of<T>(x: T) {
    println!("{}", get_type(x));
}

pub fn demo_reflect_type() {
    let x = "holy".to_string();
    let ref z = x;
    let w = &x;
    print_type_of(z);
    print_type_of(w);
}