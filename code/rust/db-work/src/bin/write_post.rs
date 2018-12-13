extern crate db_work;
extern crate diesel;

use self::db_work::*;
use std::io::{stdin, Read};

#[cfg(not(windows))]
const EOF: &'static str = "CTRL+D";

#[cfg(windows)]
const EOF: &'static str = "CTRL+Z";

fn main() {
    let connection = establish_connection();

    println!("What would you like title to be?");
    let mut title = String::new();
    stdin().read_line(&mut title).unwrap();
    let title = &title[..(title.len() - 1)];
    println!("\nOk! Let's write {} (Press {} when finished)\n", title, EOF);
    let mut body = String::new();
    stdin().read_to_string(&mut body).unwrap();

    let post = create_post(&connection, title, &body);
    println!("\nSaved draft {} with id {}", title, post.id);
}