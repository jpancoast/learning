fn main() {
    let s1 = String::from("abc");
    let s2 = s1; //This isn't a copy, this is a move.  It's the VALUE that has an OWNER

    /*
     * This DOES NOT WORK, because s2 now owns it, and not s1
     *  s2 has "borrowed" it
     *
     * https://www.udemy.com/course/ultimate-rust-crash-course/learn/lecture/17981901#overview
     */
    //println!("{}", s1)

    let s3 = String::from("blah");
    let s4 = s3.clone();

    println!("s3 = {}, s4 = {}", s3, s4);
}
