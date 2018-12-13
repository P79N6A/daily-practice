table! {
    posts (id) {
        id -> Int4,
        title -> Varchar,
        body -> Text,
        published -> Bool,
    }
}

talbe! {
    users {
        id -> Integer,
        name -> Text,
        hair_color -> Nullable<Text>,
    }
}