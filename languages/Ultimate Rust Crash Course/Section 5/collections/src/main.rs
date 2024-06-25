use std::collections::HashMap;

fn main() {
    //Vectors, this one stores i32 values
    let mut v: Vec<i32> = Vec::new();
    v.push(1);
    v.push(2);
    v.push(3);

    let x = v.pop();
    println!("{:?}", x);
    println!("{}", v[0]);

    /*
     *  the ! after things like vec! & println! mean it's a "macro"
     */
    let v2 = vec![1, 2, 3];
    println!("{:?}", v2);

    //Hashmaps, man I love these!
    let mut h: HashMap<u8, bool> = HashMap::new();

    h.insert(1, true);
    h.insert(2, false);

    println!("{:?}", h);
}
