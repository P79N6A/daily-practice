macro_rules! newtype_new {
    (struct $name:ident($t:ty);) => { newtype_new! { () struct $name($t); } };
    (pub struct $name:ident($t:ty);) => { newtype_new! { (pub) struct $name($t); } };
    
    (($($vis:tt)*) struct $name:ident($t:ty);) => {
        #[derive(Debug)]
        $($vis)* struct $name($t);
        as_item! {
            impl $name {
                $($vis)* fn new(value: $t) -> Self {
                    $name(value)
                }
            }
        }
    };
}
macro_rules! as_item { ($i:item) => {$i} }

pub fn demo_newtype_new() {
    newtype_new!(pub struct Gaga(i32););

    let g = Gaga::new(22);
    println!("{:?}", g);
}