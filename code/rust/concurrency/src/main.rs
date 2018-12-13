extern crate concurrency;

fn main() {
    // concurrency::demo_spawn_threads();
    // concurrency::demo_threads_with_closure();
    // concurrency::demo_channel();
    // concurrency::lock::demo_mutex();
    // concurrency::lock::demo_threads_join();
    // concurrency::oo::collection::demo_average_collection();

    concurrency::oo::blog::demo_post();

    concurrency::unsafety::demo_unsafe_raw_pointer_of_slice();

    concurrency::unsafety::demo_ffi();
}
