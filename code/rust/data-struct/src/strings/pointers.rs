use std::rc::Rc;

pub fn demo_pointer_addr() {
    let original = Rc::new("mazurka".to_string());
    let cloned = original.clone();
    let impostor = Rc::new("mazurka".to_string());
    
    println!("text: {}, {}, {}, pointers: {:p}, {:p}, {:p}", 
        original, cloned, impostor, original, cloned, impostor);
} 