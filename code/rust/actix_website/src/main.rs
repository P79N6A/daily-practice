extern crate actix_web;
extern crate actix_website;
use actix_web::{server, App};
use actix_website::{AppState, index};
use std::cell::Cell;

fn main() {
    server::new(|| App::with_state(AppState{ counter: Cell::new(0) })
        .resource("/", |r| r.f(index)))
        .bind("127.0.0.1:8088")
        .unwrap()
        .run();
}