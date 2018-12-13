use std::cmp::{min};

#[derive(Debug)]
struct Similarity {
    same_chars: u32,        // same chars in same position
    longest_seqs: u32,      // the longest common subsequence
}
impl Similarity {
    fn new() -> Similarity { Similarity { same_chars: 0, longest_seqs: 0 } }
}

fn similarity(a: &str, b: &str) -> Similarity {
    let mut similarity = Similarity::new();
    let mut longest_seqs = 0;
    let len = min(a.len(), b.len());
        
    for i in 0..len {
        if a.as_bytes()[i] == b.as_bytes()[i] {
            similarity.same_chars += 1;
            longest_seqs += 1;
            if similarity.longest_seqs < longest_seqs {
                similarity.longest_seqs = longest_seqs;
            }
        } else {
            longest_seqs = 0;
        }
    }

    similarity
}

pub fn demo_similarity() {
    let fort = "fort";
    let fosh = "fosh";
    let fish = "fish";
    println!("similarity between '{}' and '{}': {:?}", fosh, fort, similarity(fosh, fort));
    println!("similarity between '{}' and '{}': {:?}", fosh, fish, similarity(fosh, fish));
}