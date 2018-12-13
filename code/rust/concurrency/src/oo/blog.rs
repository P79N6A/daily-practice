pub struct Post {
    state: Option<Box<State>>,
    content: String,
}

impl Post {
    pub fn new() -> Post {
        Post {
            state: Some(Box::new(Draft {})),
            content: String::new(),
        }
    }

    pub fn add_text(&mut self, text: &str) {
        if self.state.as_ref().unwrap().get_type() == StateKind::Draft {
            self.content.push_str(text);
        }
    }

    pub fn content(&self) -> &str {
        self.state.as_ref().unwrap().content(&self)
    }

    pub fn request_review(&mut self) {
        if let Some(s) = self.state.take() {
            self.state = Some(s.request_review())
        }
    }

    pub fn approve(&mut self) {
        if let Some(s) = self.state.take() {
            self.state = Some(s.approve())
        }
    }

    pub fn reject(&mut self) {
        if let Some(s) = self.state.take() {
            self.state = Some(s.reject())
        }
    }
}

trait State {
    fn request_review(self: Box<Self>) -> Box<State>;
    fn approve(self: Box<Self>) -> Box<State>;
    fn reject(self: Box<Self>) -> Box<State>;

    fn content<'a>(&self, _post: &'a Post) -> &'a str {
        ""
    }

    fn get_type(&self) -> StateKind {
        StateKind::Unknown
    }
}

#[derive(Debug, PartialEq)]
enum StateKind {
    Draft,
    PendingReview,
    Published,
    Unknown,
}

struct Draft {}
impl State for Draft {
    fn request_review(self: Box<Self>) -> Box<State> {
        Box::new(PendingReview {})
    }

    fn approve(self: Box<Self>) -> Box<State> {
        self
    }

    fn reject(self: Box<Self>) -> Box<State> {
        self
    }

    fn get_type(&self) -> StateKind {
        StateKind::Draft
    }
}

struct PendingReview {}
impl State for PendingReview {
    fn request_review(self: Box<Self>) -> Box<State> {
        self
    }

    fn approve(self: Box<Self>) -> Box<State> {
        Box::new(Published {})
    }

    fn reject(self: Box<Self>) -> Box<State> {
        Box::new(Draft {})
    }

    fn get_type(&self) -> StateKind {
        StateKind::PendingReview
    }
}

struct Published {}
impl State for Published {
    fn request_review(self: Box<Self>) -> Box<State> {
        self
    }

    fn approve(self: Box<Self>) -> Box<State> {
        self
    }

    fn content<'a>(&self, post: &'a Post) -> &'a str {
        &post.content
    }

    fn reject(self: Box<Self>) -> Box<State> {
        self
    }

    fn get_type(&self) -> StateKind {
        StateKind::Published
    }
}

pub fn demo_post() {
    let mut post =  Post::new();
    post.add_text("hello, world");
    post.request_review();
    post.approve();
    println!("content is: {}", post.content());
}
