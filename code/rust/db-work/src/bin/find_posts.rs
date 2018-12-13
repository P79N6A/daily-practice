extern crate db_work;
extern crate diesel;

use self::db_work::*;
use self::models::*;
use self::diesel::prelude::*;

fn main() {
    use db_work::schema::posts::dsl::*;

    let connection = establish_connection();

    let t = read_line_without_crlf("what do you want to find:");

    let query = format!("%{}%", t);
    let results = posts.filter(title.ilike(query))
                        .load::<Post>(&connection)
                        .expect("Error loading posts");
    
    println!("Found {} posts contains: {}", results.len(), t);
    for post in results {
        println!("{}", "-".repeat(32));
        println!("{}\t{}", post.title, post.body);
    }                        
}