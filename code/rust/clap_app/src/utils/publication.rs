#[derive(Debug)]
pub struct Author {
    name: String,
}
impl Author {
    pub fn new(name: &str) -> Author {
        Author { name: String::from(name) }
    }
    pub fn name<'a>(&'a self) -> &'a str {
        self.name.as_str()
    }
}

#[derive(Debug)]
pub struct Post<'a> {
    title: String,
    author: &'a Author,
}
impl<'a> Post<'a> {
    pub fn new(title: &str, author: &'a Author) -> Post<'a> {
        Post {
            title: String::from(title),
            author,
        }
    }
    pub fn author(&self) -> Option<&Author> {
        Some(self.author)
    }
}

pub fn demo_publication() {
    let author = Author::new("Adam");
    let post = Post::new("cookbook", &author);

    let author_name = match post.author() {
        Some(author) => author.name(),
        None => "No Author",
    };

    println!("author: {}", author_name);
}

