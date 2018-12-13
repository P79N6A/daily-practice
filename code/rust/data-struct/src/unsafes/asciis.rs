use std::ascii::AsciiExt;

#[derive(Debug, Eq, PartialEq)]
struct NotAsciiError(Vec<u8>);

#[derive(Debug, Eq, PartialEq)]
struct Ascii (Vec<u8>);

impl Ascii {
    fn from_bytes(bytes: Vec<u8>) -> Result<Ascii, NotAsciiError> {
        if bytes.iter().any(|&byte| !byte.is_ascii()) {
            return Err(NotAsciiError(bytes))
        }
        Ok(Ascii(bytes))
    }

    unsafe fn from_bytes_unchecked(bytes: Vec<u8>) -> Ascii {
        Ascii(bytes)
    }
}

impl From<Ascii> for String {
    fn from(ascii: Ascii) -> String {
        unsafe { String::from_utf8_unchecked(ascii.0) }
    }
}

pub fn demo_ascii_within_unsafe() {
    let bytes: Vec<u8> = b"ASCII and ye shall receive".to_vec();

    let ascii: Ascii = Ascii::from_bytes(bytes.clone()).unwrap();

    let string = String::from(ascii);
    
    println!("{:?} to_string is:{}", bytes, string);
}