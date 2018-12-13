fn binary_search(list: &mut [i32], item: i32) -> Option<usize> {
    list.sort();
    let mut low = 0;
    let mut high = list.len() - 1;

    while low <= high {
        let mid = (low + high) / 2;
        let guess = list[mid];

        if guess == item {
            return Some(mid);
        } else if guess > item {
            high = mid - 1;
        } else if guess < item {
            low = mid + 1;
        }
    }
    None
}

pub fn demo_binary_search() {
    let mut list = (1..66).collect::<Vec<i32>>();
    let result = binary_search(&mut list, 32);

    println!("Binary Search Result: {:?}", result);
}
