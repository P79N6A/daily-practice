extern crate rocket;

use std::path::PathBuf;
use rocket_contrib::Json;
use rocket::http::{Cookie, Cookies};
use http::entities::Person;

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[get("/test")]
fn test() -> &'static str {
    "test!"
}

#[get("/hello/<name>/<age>/<cool>")]
fn hello(name: String, age: u8, cool: bool) -> String {
    match cool {
        true => format!("You're a cool {} year old, {}", age, name),
        false => format!("{}, we need to talk about your coolness.", name)
    }
}

#[get("/holy/<paths..>")]
fn holy(paths: PathBuf) -> String {
    format!("Current path is: holy/{}", paths.to_str().unwrap())
}

#[get("/json")]
fn json<'a>() -> Json<Person<'a>> {
    Json(Person {
        name: "WTF",
        sex: "male",
        age: 666,
    })
}

#[get("/cookie")]
fn cookie(mut cookies: Cookies) -> Option<String> {
    cookies.add(Cookie::new("user", "moye"));
    cookies.get("user")
        .map(|c| format!("user: {}", c.value()))
}

#[get("/heap")]
fn heap() -> String {
    let point = Box::new((0.625, "PI"));
    format!("{:?}", point)
}

pub fn launch() {
    rocket::ignite().mount("/", routes![index, test, hello, holy, json, cookie, heap]).launch();
}
