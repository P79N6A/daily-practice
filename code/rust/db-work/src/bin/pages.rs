extern crate db_work;
extern crate diesel;

use self::db_work::*;
use self::models::*;
use self::diesel::prelude::*;

fn main() {
    use db_work::schema::*;
    use db_work::pagination::*;

    let connection = establish_connection();

    let query = posts::table
        .select(posts::all_columns)
        .paginate(1)
        .per_page(10);

    let (results, total_pages) = query.load_and_count_pages::<(Post)>(&connection).unwrap();

    println!("Found {} posts.", results.len());
    for post in results {
        println!("{}", "-".repeat(32));
        println!("{}\t{}", post.title, post.body);
    }
}