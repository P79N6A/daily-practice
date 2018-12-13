extern crate minigrep;

use std::env;
use std::process;

use minigrep::Config;

fn main() {
    let config = Config::new(env::args()).unwrap_or_else(|err| {
        eprintln!("Problem parsing arguments: {} by \"{}\"", err, args.join (" "));
        process::exit(1);
    });

    println!("Searching for \"{}\", In file {}:", config.query, config.filename);

    minigrep::run(config).unwrap_or_else(|err| {
        eprintln!("Application error: {}", err);
        process::exit(1);
    });
}
