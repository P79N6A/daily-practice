use utils::println;

fn quick_sort(list: &[i32]) -> Vec<i32> {
    if list.len() < 2 {
        list.to_owned()
    } else {
        let pivot = list[0];
        let (less, greater): (Vec<i32>, Vec<i32>) =
            list.iter().cloned().skip(1).partition(|&x| x <= pivot);

        let mut result = quick_sort(&less);
        result.push(pivot);
        result.extend(quick_sort(&greater));
        result
    }
}

pub fn demo_quick_sort() {
    let v = vec![6, 8, 1, 5, 12, 4, 3, 7];
    let result = quick_sort(&v);

    println(result);
}
