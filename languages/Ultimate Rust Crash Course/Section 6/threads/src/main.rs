use std::thread;

fn main() {
    let handle = thread::spawn(move || {
        //do stuff in child thread
        println!("Hello from child thread!");
    });

    println!("Hello from main thread");

    //wait until thread has exited
    handle.join().unwrap();
}
