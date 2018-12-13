use std::panic::catch_unwind;
use std::error::Error;
use std::fmt;
use std::io::{Write, stderr, Error as IoError, ErrorKind};

pub fn catch_fault() {
    let result = catch_unwind(|| {
        panic!("holy crap!");
    });

    println!("\n{:?}\n", result);
}

pub fn get_something() -> Result<(), String> {
    Err("holy crap".to_string())
}

pub fn print_error(mut err: &Error) {
    let _ = writeln!(stderr(), "error: {}", err);
    while let Some(cause) = err.cause() {
        let _ = writeln!(stderr(), "caused by: {}", cause);
        err = cause;
    }
}

pub fn demo_print_std_err() {
    #[derive(Debug)]
    struct SuperErrorSideKick;
    impl fmt::Display for SuperErrorSideKick {
        fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
            write!(f, "SuperErrorSideKick is here!")
        }
    }
    impl Error for SuperErrorSideKick {
        fn description(&self) -> &str {
            "I'm SuperError side kick"
        }
    }
    #[derive(Debug)]
    struct SuperError;
    impl fmt::Display for SuperError {
        fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
            write!(f, "SuperError is here!")
        }
    }
    impl Error for SuperError {
        fn description(&self) -> &str {
            "I'm the superhero of errors"
        }
        fn cause(&self) -> Option<&Error> {
            Some(&SuperErrorSideKick{})
        }
    }

    let mut err = SuperError{};
    println!("{:?}", err.cause());
    print_error(&err);
}

pub fn err_propagation() -> Result<(), IoError> {
    Err(IoError::new(ErrorKind::Other, "oh no!"))
}

pub fn demo_err_propagation() -> Result<(), IoError> {
    let result = err_propagation()?;

    println!("propagation in middle.");

    Ok(())
}