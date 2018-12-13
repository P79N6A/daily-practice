#[derive(Debug)]
struct Account {
    name: String,
    language: String,
    sex: String,
    age: u8,
}

pub fn run() {
    let peter = Account {
        name: "Peter".to_string(),
        language: "English".to_string(),
        sex: "Male".to_string(),
        age: 22,
    };

    match peter {
        // ref members
        Account {
            ref name,
            ref language,
            age,
            ..
        } if age < 20 =>
        {
            println!("age < 20, {}, {}", name, language)
        }
        Account {
            ref name,
            ref language,
            ..
        } => println!("{}, {}", name, language),
    }

    println!("{:?}", peter);
}
