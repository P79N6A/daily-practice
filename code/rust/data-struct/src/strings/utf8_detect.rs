use utils::println;

fn detect_utf8(text: &str) {
    println(text);
    let bytes = text.as_bytes();
    for c in bytes {
        print!("{:b} ", c);
        match c {
            d if d >> 7 == 0b0u8 => print!(" UTF8 √\n"),
            e if e >> 6 == 0b10u8 => print!(" UTF8 √\n"),
            f if f >> 5 == 0b110u8 => print!(" UTF8 √\n"),
            g if g >> 4 == 0b1110u8 => print!(" UTF8 √\n"),
            h if h >> 3 == 0b11110u8 => print!(" UTF8 √\n"),
            _ => print!(" is not an UTF8 char.\n"),
        }
    }
}

pub fn demo_detect_utf8() {
    detect_utf8("うどん");
}
