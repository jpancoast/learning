fn main() {
    let num: i32 = 5;

    let msg = if num == 5 {
        "five"
    } else if num == 4 {
        "four"
    } else {
        "other"
    };

    println!("{}", msg);

    'bob: loop {
        loop {
            loop {
                break 'bob;
            }
        }
    }

    for num in [7, 8, 9].iter() {
        println!("{}", num);
    }

    let array = [(1, 2), (3, 4)];

    for (x, y) in array.iter() {
        println!("{}, {}", x, y);
    }

    for num in 0..5 {
        println!("{}", num);
    }

    println!();

    //The '=' makes it inclusive where it's normally exclusive.
    for num in 0..=5 {
        println!("{}", num);
    }
}
