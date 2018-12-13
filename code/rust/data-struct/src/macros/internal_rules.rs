macro_rules! foo {
    (@expr_rule $e:expr) => { 
        println!("holy expr! {}", stringify!($e) )
    };

    ($($tts:tt)*) => {
        foo!(@expr_rule $($tts)*)
    };
}

pub fn demo_internal_ruels() {
    foo!(3+5+9 != 17 ?);
}
