#[macro_use]
extern crate data_struct;

use data_struct::{algorithms, closures, iterators, operators, strings, structs, 
                traits, IO, concurrency, low_level, macros, unsafes, ffi};

fn main() {
    structs::tree::run();

    structs::match_ref::run();

    traits::generics::demo_generics();

    traits::associated_type::demo_associated_type();

    traits::dot::demo_dots();

    operators::complex::demo_complexes();

    operators::binary::demo_binary();

    traits::appellation::demo_appellation();

    traits::borrowable::demo_borrowable();

    traits::cows::demo_cow();

    closures::sorting::demo_sorting();

    closures::drops::demo_drops();

    iterators::dump::demo_dump();

    iterators::scan_fib::demo_scan_fib_vec();

    algorithms::binary::demo_binary_search();

    algorithms::quick_sort::demo_quick_sort();

    iterators::fizzbuzzs::demo_fizzbuzzs();

    iterators::i32range::demo_i32range();

    strings::utf8_detect::demo_detect_utf8();

    algorithms::bfs::demo_bfs_search();

    algorithms::dijkstra::demo_dijkstra();

    algorithms::dynamic_programming::demo_similarity();

    strings::cow::demo_cow();

    strings::pointers::demo_pointer_addr();

    strings::regexs::demo_regex();

    IO::grep::demo_grep();

    IO::collector::demo_collect_lines();

    IO::read_dir::demo_read_dir();

    // IO::echo::serve("127.0.0.1:9999").expect("error: ");

    IO::client::demo_client();

    closures::high_order::demo_high_order();

    concurrency::tasks_queue::demo_tasks();    

    low_level::reflection::demo_reflect_type();

    // concurrency::interaction::demo_interaction();

    // concurrency::mutexes::demo_mutexes();

    concurrency::rwlocks::demo_rwlocks();

    concurrency::condvars::demo_condvar();

    concurrency::atomics::demo_atomics();

    let v = vec_new![1, 2, 3];
    let a = json!({"name" : "Jack", "vector": v});
    let q = stringify!(json!({"name" : "Jack", "vector": v}));
    println_more!("{} = {:?}\n", q, a);

    macros::including::demo_include();

    unsafes::asciis::demo_ascii_within_unsafe();

    unsafes::ref_with_flag::demo_ref_with_flag();

    // ffi::libgit::demo_libgit();

    macros::diff_expr_tt::demo_difference_of_expr_and_tt();
    macros::ident_self::demo_ident_self();

    macros::newtype_new::demo_newtype_new();
    macros::internal_rules::demo_internal_ruels();

    closures::impl_traits::demo_impl_traits();

    traits::generators::demo_generator();
}
