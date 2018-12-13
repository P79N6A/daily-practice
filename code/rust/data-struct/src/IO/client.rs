extern crate reqwest;

use self::reqwest::{Client, Result, Response};

fn request(target: &str) -> Result<Response> {
    let client = Client::new();
    client.get(target).send()
}

pub fn demo_client() {
    let target = "http://example.org/";
    println!("{} RESP: {}", target, request(target).unwrap().status());
}