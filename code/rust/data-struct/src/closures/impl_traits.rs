use std::str::FromStr;

// universally quantified types:
//      the caller, provide the type you want, T, and then the function returns it. 
fn parse<F>(s: &str) -> Result<F, <F as FromStr>::Err> 
    where F: FromStr
{
    s.parse::<F>()
}
// existential types:
//      the caller can't choose, and the function itself gets to choose.
fn parse2(s: &str) -> impl FromStr {
    let t:u32 = s.parse().unwrap();
    t
}

pub fn demo_impl_traits() {
    let i:i32 = parse("32").unwrap();
    println!("universal returned type: {}", i);

    let _ = parse2("64"); 
}

// trait Animal {}

// struct Bird {}
// struct Horse {}

// impl Animal for Bird {}
// impl Animal for Horse {}

// fn foo1(x: &impl Animal, y: &impl Animal) {}
// fn foo2<T1: Animal, T2: Animal>(x: &T1, y: &T2) {}
// fn bar<T: Animal>(x: &T, y: &T) {}