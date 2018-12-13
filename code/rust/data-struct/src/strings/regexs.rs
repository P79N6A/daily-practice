extern crate regex;

use self::regex::Regex;

pub fn demo_regex() {
    let semver = Regex::new(r"(\d+)\.(\d+)\.(\d+)(-[-.[:alnum:]]*)?").unwrap();
    let haystack = r#"regex = "0.2.5""#;

    let captures = semver.captures(haystack)
                    .ok_or("semver regex should have matched").unwrap();

    println!("regex captures: {:?}", captures);
}