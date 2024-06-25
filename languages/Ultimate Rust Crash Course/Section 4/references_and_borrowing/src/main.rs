fn main() {
    let mut s1 = String::from("abc");
    do_stuff(&mut s1);

    println!("{}", s1);
}

fn do_stuff(s: &mut String) {
    /*
     * The '.' in s.insert_str automatically derefs the reference, this also works
     * (*s).insert_str(0, "Hi, ")
     */
    s.insert_str(0, "Hi, ");
    (*s).insert_str(0, "Hi, ");
}
