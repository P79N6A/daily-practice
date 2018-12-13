# Ownership Rules

First, let’s take a look at the ownership rules. Keep these rules in mind as we work through the examples that illustrate the rules:

1. Each value in Rust has a variable that’s called its owner.
2. There can only be one owner at a time.
3. When the owner goes out of scope, the value will be dropped.

# Lifetime Elision

1. Each parameter that is a reference gets its own lifetime parameter. In other words, a function with one parameter gets one lifetime parameter: fn foo<'a>(x: &'a i32), a function with two arguments gets two separate lifetime parameters: fn foo<'a, 'b>(x: &'a i32, y: &'b i32), and so on.
2. If there is exactly one input lifetime parameter, that lifetime is assigned to all output lifetime parameters: fn foo<'a>(x: &'a i32) -> &'a i32.
3. If there are multiple input lifetime parameters, but one of them is &self or &mut self because this is a method, then the lifetime of self is assigned to all output lifetime parameters. This makes writing methods much nicer.

# Box<T>, Rc<T>, or RefCell<T>:

1. ```Rc<T>``` enables multiple owners of the same data; ```Box<T>``` and ```RefCell<T>``` have single owners.
2. ```Box<T>``` allows immutable or mutable borrows checked at compile time; ```Rc<T>``` allows only immutable borrows checked at compile time; ```RefCell<T>``` allows immutable or mutable borrows checked at runtime.
3. Because ```RefCell<T>``` allows mutable borrows checked at runtime, you can mutate the value inside the ```RefCell<T>``` even when the ```RefCell<T>``` is immutable.