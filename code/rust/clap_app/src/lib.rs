pub mod utils;
pub mod asyncio;

extern crate clap;
extern crate tokio;

use clap::{Arg, App};
use utils::{ demo_shadowing, show, Table, StringTable, publication, fault };
use asyncio::tokio_server;

pub fn entry() {
    let matches = App::new("holy-clap")
                        .version("1.0.0")
                        .about("Do something")
                        .author("Mo Ye")
                        .arg(Arg::with_name("v")
                            .short("v")
                            .multiple(true)
                            .help("Sets the level of verbosity"))
                        .get_matches();

    match matches.occurrences_of("v") {
        0 => println!("No verbose info"),
        1 => println!("Some verbose info"),
        2 => println!("Tons of verbose info"),
        3 | _ => println!("Don't be crazy"),
    }

    let mut table = Table::new();
    table.insert("Gesualdo".to_string(),
        vec!["many madrigals".to_string(),
        "Tenebrae Responsoria".to_string()]);
    table.insert("Caravaggio".to_string(),
        vec!["The Musicians".to_string(),
        "The Calling of St. Matthew".to_string()]);
    table.insert("Cellini".to_string(),
        vec!["Perseus with the head of Medusa".to_string(),
        "a salt cellar".to_string()]);
    show(&table);

    demo_shadowing();

    StringTable::demo_finding();

    publication::demo_publication();

    // fault::catch_fault();

    // tokio_server::serve("127.0.0.1:12345");    

    let result = fault::get_something();
    if let Some(err) = result.err() {
        println!("{:?}", err);
    }

    fault::demo_print_std_err();

    let result = fault::demo_err_propagation();
    println!("{:?}", result);

    utils::structs::demo_structs();

    utils::refs::demo_refs();
}