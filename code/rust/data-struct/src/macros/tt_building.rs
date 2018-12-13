macro_rules! call_a_or_b_on_tail {
    ((a: $a:expr, b: $b:expr), call a: $($tail:tt)*) => {
        $a(stringify!($($tail)*))
    };

    ((a: $a:expr, b: $b:expr), call b: $($tail:tt)*) => {
        $b(stringify!($($tail)*))
    };

    ($ab:tt, $_skip:tt $($tail:tt)*) => {
        call_a_or_b_on_tail!($ab, $($tail)*)
    };
}

fn compute_len(s: &str) -> Option<usize> {
    Some(s.len())
}

fn show_tail(s: &str) -> Option<usize> {
    println!("tail: {:?}", s);
    None
}

pub fn demo_tt_building() {
    assert_eq!(
        call_a_or_b_on_tail!(
            (a: compute_len, b: show_tail),
            the recursive part that skips over all these
            tokens doesn't much care whether we will call a
            or call b: only the terminal rules care.
        ),
        None
    );
    assert_eq!(
        call_a_or_b_on_tail!(
            (a: compute_len, b: show_tail),
            and now, to justify the existence of two paths
            we will also call a: its input should somehow
            be self-referential, so let's make it return
            some ninety one!
        ),
        Some(91)
    );
}