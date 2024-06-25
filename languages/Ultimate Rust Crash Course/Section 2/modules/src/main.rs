use modules::greet;

fn main() {
    greet();

    let x: i32 = rand::random();
    println!("{}", x);
}
