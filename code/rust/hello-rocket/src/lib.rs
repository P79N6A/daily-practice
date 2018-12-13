#![feature(plugin)]
#![plugin(rocket_codegen)]

#[macro_use] extern crate serde_derive;
#[macro_use] extern crate rocket_contrib;

extern crate rocket;

pub mod http;
