use std::borrow::Cow;
use std::path::PathBuf;

#[derive(Debug)]
enum Error {
    OutOfMemory,
    StackOverflow,
    MachineOnFire,
    Unfathomable,
    FileNotFound(PathBuf),
}

fn describe(error: &Error) -> Cow<'static, str> {
    match *error {
        Error::OutOfMemory => "out of memory".into(),
        Error::StackOverflow => "stack overflow".into(),
        Error::MachineOnFire => "machine on fire".into(),
        Error::Unfathomable => "machine bewildered".into(),
        Error::FileNotFound(ref path) => format!("file not found: {}", path.display()).into(),
    }
}

pub fn demo_cow() {
    let error = Error::FileNotFound(PathBuf::from("/etc/hosts"));

    let mut log = Vec::<String>::new();
    log.push(describe(&error).into_owned());

    println!("{:?}", error);
    println!("{:?}", log);
}
