#![feature(proc_macro, wasm_custom_section, wasm_import_module)]
extern crate wasm_bindgen;
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

#[wasm_bindgen(module = "./index")]
extern {
    fn duang(name: String);
}

#[wasm_bindgen]
pub fn greet(name: &str) {
    duang(format!("Hello, {}", name))
}