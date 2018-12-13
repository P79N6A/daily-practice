extern crate actix;
extern crate futures;

pub mod apps;

pub fn run() {
    apps::app::run();
}