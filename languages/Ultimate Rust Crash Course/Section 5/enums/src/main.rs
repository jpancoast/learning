use std::fs::File;
use DispenserItem::*;

fn main() {
    let thing_1 = Empty;
    let thing_2 = Ammo(10);
    let thing_3 = Things("thing".to_string(), 5);
    let thing_4 = Place { x: 10, y: 10 };

    thing_1.display();
    println!();

    thing_2.display();
    println!();

    thing_3.display();
    println!();

    thing_4.display();

    /*
     * Option<T> is an Enum that has two variants: Some(T) and None,
     *  and is part of the standard rust library. is_some & is_none are
     *  functions defined on Option<T> that return a bool.
     */
    let mut x: Option<i32> = None;

    x = Some(5);

    println!("{}", x.is_some());
    println!("{}", x.is_none());

    for i in x {
        println!("{}", i);
    }

    /*
     * Result<T, E> is an Enum that has two variants: Ok(T) and Err(E),
     */
    let res = File::open("foo");

    match res {
        Ok(f) => println!("{:?}", f),
        Err(e) => println!("{}", e),
    }
}

enum DispenserItem {
    Empty,
    Ammo(u8),
    Things(String, i32),
    Place { x: i32, y: i32 },
}

impl DispenserItem {
    fn display(&self) {
        match self {
            Empty => println!("empty"),
            Ammo(x) => println!("ammo: {}", x),
            Things(x, y) => println!("thing: {} {}", x, y),
            Place { x, y } => println!("place: {} {}", x, y),
        }
    }
}
