use std::collections::HashMap;
use std::convert::From;

#[derive(Clone, PartialEq, Debug)]
pub enum Json {
    Null,
    Boolean(bool),
    Number(f64),
    String(String),
    Array(Vec<Json>),
    Object(Box<HashMap<String, Json>>),
}
impl From<bool> for Json {
    fn from(t: bool) -> Json { Json::Boolean(t) }
}
impl From<String> for Json {
    fn from(t: String) -> Json { Json::String(t) }
}
impl<'a> From<&'a str> for Json {
    fn from(t: &'a str) -> Json { Json::String(t.to_string()) }
}
impl<'a> From<Vec<&'a str>> for Json {
    fn from(t: Vec<&'a str>) -> Json {
        Json::Array(t.into_iter().map(|a| Json::from(a)).collect())
    }
}
macro_rules! impl_from_num_for_json {
    ( $( $t:ident )* ) => {
        $(
            impl From<$t> for Json {
                fn from(t: $t) -> Json { Json::Number(t as f64) }
            }
        )*
    };
}
impl_from_num_for_json!(u8 i8 u16 i16 u32 i32 u64 i64 usize isize f32 f64);

macro_rules! impl_from_num_for_vec {
    ( $( $t:ident )* ) => {
        $(
            impl From<Vec<$t>> for Json {
                fn from(t: Vec<$t>) -> Json {
                    Json::Array(t.into_iter().map(|a| Json::from(a)).collect())
                }
            }
        )*
    };
}
impl_from_num_for_vec!(u8 i8 u16 i16 u32 i32 u64 i64 usize isize f32 f64 bool String);

#[macro_export]
macro_rules! json {
    (null) => {
        macros::jsons::Json::Null
    };
    ([ $( $element:expr ),* ]) => {
        macros::jsons::Json::Array(vec![ $( json!($element) ),* ])
    };
    ({ $( $key:tt : $value:tt ),* }) => {
        {
            let mut fields = Box::new(std::collections::HashMap::new());
            $( fields.insert($key.to_string(), json!($value)); )*
            macros::jsons::Json::Object(fields)
        }
    };
    ($other:tt) => {
        macros::jsons::Json::from($other)
    };
}