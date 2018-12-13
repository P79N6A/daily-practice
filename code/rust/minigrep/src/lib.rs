use std::env;
use std::error::Error;
use std::fs::File;
use std::io::prelude::*;

pub struct Config {
    pub query: String,
    pub filename: String,
    pub case_sensitive: bool,
}

impl Config {
    pub fn new(args: std::env::Args) -> Result<Config, &'static str> {
        args.next();

        let query = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a query string"),
        };

        let filename = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a file name"),
        };

        let case_sensitive = env::var("CASE_INSENSITIVE").is_err();
        // if !case_sensitive {
        //     println!("case_sensitive::::{}", env::var("CASE_INSENSITIVE").unwrap());
        // }

        Ok(Config { query, filename, case_sensitive })
    }
}

pub fn run(config: Config) -> Result<(), Box<Error>>{
    let mut contents = String::new();
    File::open(config.filename)?.read_to_string(&mut contents)?;

    for line in search(&config.query, &contents, config.case_sensitive) {
        println!("{}", line);
    }

    Ok(())
}

pub fn search<'a>(query: &str, contents: &'a str, case_sensitive: bool) -> Vec<&'a str> {
    let mut results = Vec::new();

    for line in contents.lines() {
        if !case_sensitive {
            let query = query.to_lowercase();
            if line.to_lowercase().contains(&query) {
                results.push(line)
            }
        } else {
            if line.contains(query) {
                results.push(line)
            }
        }
    }

    results
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn case_sensitive() {
        let query = "duct";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.
Duct tape.";

        assert_eq!(
            vec!["safe, fast, productive."],
            search(query, contents, true)
        );
    }

    #[test]
    fn case_insensitive() {
         let query = "rUsT";
         let contents = "\
Rust:
safe, fast, productive.
Pick three.
Trust me.";

         assert_eq!(
             vec!["Rust:", "Trust me."],
             search(query, contents, false)
         );
    }
}
