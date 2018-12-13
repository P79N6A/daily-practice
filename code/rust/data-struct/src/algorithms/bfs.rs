use std::rc::Rc;
use std::collections::VecDeque;

#[derive(Debug, Clone)]
struct Item {
    neighbours: Option<Vec<Rc<Item>>>,
    value: String,
}
impl Item {
    fn new(v: &str) -> Item {
        Item {
            neighbours: None,
            value: v.to_string(),
        }
    }
    fn link_to(&mut self, neighbour: Item) {
        let mut vec = Vec::new();
        if let Some(original) = self.neighbours.take() {
            vec = original.to_vec();
        }
        vec.push(Rc::new(neighbour));
        self.neighbours = Some(vec);
    }
}

#[derive(Debug)]
struct SearchResult {
    success: bool,
    depth: Option<u32>,
}

fn bfs(root: Item, target: &str) -> SearchResult {
    let mut queue = VecDeque::new();

    if let Some(neighbours) = root.neighbours {
        for item in neighbours {
            queue.push_back(item);
        }
        return bfs_search(queue, target, 1);
    }

    SearchResult {
        success: false,
        depth: None,
    }
}

fn bfs_search(mut queue: VecDeque<Rc<Item>>, target: &str, round: u32) -> SearchResult {
    let mut current_round = VecDeque::new();

    while !queue.is_empty() {
        if let Some(item) = queue.pop_front() {
            if item.value == target {
                return SearchResult {
                    success: true,
                    depth: Some(round),
                };
            }
            if let Some(neighbours) = item.neighbours.as_ref() {
                for item in neighbours.clone() {
                    current_round.push_back(item);
                }
            }
        }
    }

    bfs_search(current_round, target, round + 1)
}

pub fn demo_bfs_search() {
    let mut root = Item::new("You");
    let mut jack = Item::new("Jack");
    let mut rose = Item::new("Rose");
    let lucy = Item::new("Lucy");
    let lily = Item::new("Lily");
    jack.link_to(lucy);
    rose.link_to(lily);
    root.link_to(jack);
    root.link_to(rose);

    let result = bfs(root, "Lucy");
    println!("BFS search result: {:?}", result);
}
