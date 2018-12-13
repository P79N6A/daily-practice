extern crate actix_web;
use actix_web::{HttpRequest};
use std::cell::Cell;

pub struct AppState {
    pub counter: Cell<usize>,
}

pub fn index(req: &HttpRequest<AppState>) -> String {
    let count = req.state().counter.get() + 1;
    req.state().counter.set(count);

    format!("Counter: {}", count)
}