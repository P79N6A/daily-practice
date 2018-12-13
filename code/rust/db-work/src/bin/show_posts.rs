extern crate db_work;
extern crate diesel;

use self::db_work::*;
use self::models::*;
use self::diesel::prelude::*;

fn main() {
    use db_work::schema::posts::dsl::*;

    let connection = establish_connection();
    let results = posts.filter(published.eq(true))
                        .limit(5)
                        .load::<Post>(&connection)
                        .expect("Error loading posts");

    println!("Displaying {} posts", results.len());
    for post in results {
        println!("{}", "-".repeat(32));
        println!("{}\t{}", post.title, post.body);
    }
}