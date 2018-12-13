#[macro_use]
extern crate diesel;
extern crate dotenv;

pub mod schema;
pub mod models;
pub mod pagination;

use diesel::prelude::*;
use diesel::pg::PgConnection;
use dotenv::dotenv;
use self::models::{Post, NewPost};
use std::env;
use std::io::{stdin};
use std::borrow::Cow;

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL")
        .expect("DATABASE_URL must be set");
    PgConnection::establish(&database_url)
        .expect(&format!("Error connecting to {}", database_url))
}

pub fn create_post<'a>(conn: &PgConnection, 
                        title: &'a str, 
                        body: &'a str) -> Post 
{
    use schema::posts;

    let new_post = NewPost {
        title,
        body,
    };

    diesel::insert_into(posts::table)
        .values(&new_post)
        .get_result(conn)
        .expect("Error saving new post")
}

pub fn read_line_without_crlf<'a>(prompt: &'static str) 
    -> Cow<'a, str> 
{ // trim all of \n\t\s
    println!("{}", prompt);
    let mut t = String::new();
    stdin().read_line(&mut t).unwrap();
    Cow::Owned(t.trim().to_owned())
}