fn main() {
    let add = |x, y| x + y;

    let res = add(1, 2);
    println!("{}", res);

    let blah = "sldfjalsdjf".to_string();
    let f = || {
        println!("{}", blah);
    };

    f();
}
