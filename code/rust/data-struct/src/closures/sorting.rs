use std::thread;

#[derive(Debug)]
struct City {
    name: String,
    population: i64,
}
impl City {
    fn new(name: &str, population: i64) -> City {
        City {
            name: name.to_owned(),
            population,
        }
    }
}

fn start_sorting_thread(mut cities: Vec<City>) -> thread::JoinHandle<Vec<City>> {
    let key_fn = move |city: &City| -> i64 { -city.population };

    thread::spawn(move || {
        cities.sort_by_key(key_fn);
        cities
    })  
}

pub fn demo_sorting() {
    let mut cities = Vec::<City>::new();
    cities.push(City::new("New York", 300));
    cities.push(City::new("Tokyo", 120));
    cities.push(City::new("Peking", 600));

    let cities = start_sorting_thread(cities).join();
    println!("{:?}", cities);
}
