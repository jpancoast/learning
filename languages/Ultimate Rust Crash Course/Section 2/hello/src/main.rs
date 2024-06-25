/*
 * Variables are immutable in rust
 */
fn main() {
    let mut bunnies = 2 //Rust infers the type, this is mutable
    let bunnies_thing :i32 = 2 //i32 is the type

    bunnies = 3; //error if I hadn't used 'mute' (vars are immutable by default)
    println!("Hello, world!");
}
