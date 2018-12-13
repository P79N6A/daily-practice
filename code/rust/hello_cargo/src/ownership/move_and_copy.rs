pub fn demo() {
    // copy
    let s1 = "hello";
    let s2 = s1;
    println!("{}", s1);

    // move
    let s3 = String::from("world");
    let s4 = s3;
    println!("{}", s3);
}
