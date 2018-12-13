use std::collections::{HashMap, HashSet};
use std::collections::hash_map::{self, Iter};
use std::f32::INFINITY;

#[derive(Debug, Eq, PartialEq, Hash, Clone)]
struct Node {
    title: String,
    is_start: bool,
    is_end: bool,
}
impl Node {
    fn new(t: &str, is_start: bool, is_end: bool) -> Node {
        Node {
            title: t.to_owned(),
            is_start,
            is_end,
        }
    }
}

#[derive(Debug)]
struct Graph {
    graph: HashMap<Node, HashMap<Node, f32>>,
}
impl Graph {
    fn new() -> Graph {
        Graph {
            graph: HashMap::new(),
        }
    }
    fn set(&mut self, parent: &Node, child: &Node, distance: f32) {
        let mut children = HashMap::new();
        if let Some(children_ref) = self.graph.get(parent) {
            children = children_ref.clone();
        }
        children.insert(child.clone(), distance);
        self.graph.insert(parent.clone(), children);
    }
    fn set_empty(&mut self, parent: &Node) {
        self.graph.insert(parent.clone(), HashMap::new());
    }
    fn keys(&self) -> hash_map::Keys<Node, HashMap<Node, f32>> {
        self.graph.keys()
    }
    fn get(&self, n: &Node) -> Option<HashMap<Node, f32>> {
        match self.graph.get(n) {
            Some(hs) => Some(hs.clone()),
            None => None,
        }
    }
}

#[derive(Debug)]
struct Costs {
    costs: HashMap<Node, f32>,
}
impl Costs {
    fn new(graph: &Graph) -> Costs {
        let mut costs = Costs {
            costs: HashMap::new(),
        };
        if let Some(start) = graph.keys().find(|node| node.is_start) {
            if let Some(dests) = graph.get(start) {
                for (node, cost) in dests.clone() {
                    costs.set(&node, cost);
                }
            }
        }
        if let Some(end) = graph.keys().find(|node| node.is_end) {
            costs.set(&end, INFINITY);
        }
        costs
    }
    fn set(&mut self, node: &Node, cost: f32) {
        self.costs.insert(node.clone(), cost);
    }
    fn get(&self, node: &Node) -> f32 {
        match self.costs.get(node).cloned() {
            Some(cost) => cost,
            None => INFINITY,
        }
    }
    fn iter(&self) -> Iter<Node, f32> {
        self.costs.iter()
    }
}

struct Processed {
    processed: HashSet<Node>,
}
impl Processed {
    fn new() -> Processed {
        Processed {
            processed: HashSet::new(),
        }
    }
    fn exists(&self, node: &Node) -> bool {
        self.processed.contains(node)
    }
    fn set(&mut self, node: &Node) {
        self.processed.insert(node.clone());
    }
}

struct Parents {
    parents: HashMap<Node, Node>,
}
impl Parents {
    fn new() -> Parents {
        Parents {
            parents: HashMap::new(),
        }
    }
    fn set(&mut self, child: &Node, parent: &Node) {
        self.parents.insert(child.clone(), parent.clone());
    }
}

fn least_cost_node(processed: &Processed, costs: &Costs) -> Option<Node> {
    let mut least_cost = INFINITY;
    let mut least_cost_node = None;
    for (node, cost) in costs.iter() {
        if *cost < least_cost && !processed.exists(node) {
            least_cost = *cost;
            least_cost_node = Some(node.clone());
        }
    }
    least_cost_node
}

fn dijkstra(graph: &Graph) -> (Costs, Parents) {
    let mut costs = Costs::new(&graph);
    let mut processed = Processed::new();
    let mut parents = Parents::new();

    while let Some(node) = least_cost_node(&processed, &costs) {
        let cost = costs.get(&node);
        if let Some(neighbours) = graph.get(&node) {
            for (n, distance) in neighbours {
                let cost_new = cost + distance;
                if costs.get(&n) > cost_new {
                    costs.set(&n, cost_new);
                    parents.set(&n, &node);
                }
            }
        }
        processed.set(&node);
    }

    (costs, parents)
}

pub fn demo_dijkstra() {
    let start = Node::new("Start", true, false);
    let a = Node::new("A", false, false);
    let b = Node::new("B", false, false);
    let end = Node::new("End", false, true);

    let mut graph = Graph::new();
    graph.set(&start, &a, 6f32);
    graph.set(&start, &b, 2f32);
    graph.set(&a, &end, 1f32);
    graph.set(&b, &a, 3f32);
    graph.set(&b, &end, 5f32);
    graph.set_empty(&end);

    let (costs, parents) = dijkstra(&graph);
    println!("{:?}", costs);
}
