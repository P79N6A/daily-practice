pub fn from_strs(elements: Vec<&str>) -> Vec<String> {
    let mut strings = Vec::new();
    for ele in elements { 
        strings.push(String::from(ele));
    }
    strings
}