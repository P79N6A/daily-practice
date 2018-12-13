

pub fn demo_include() {
    let emoj_string = include!("emoj.in");
    println_more!("[include!(EMOJ)] = {}\n", emoj_string);
}