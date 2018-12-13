use std::borrow::Cow;
use std::env;

fn get_name() -> Cow<'static, str> {
    env::var("USER")
        .map(|v| Cow::Owned(v))
        .unwrap_or(Cow::Borrowed("whoever you are"))
}

pub fn demo_cow() {
    let mut name = get_name();
    name.to_mut().insert_str(0, "hello, ");
    println!("COW: {}", name);

    let name2 = get_name();
    name = name2.to_owned();
    println!("COW owned str: {}", name);
}