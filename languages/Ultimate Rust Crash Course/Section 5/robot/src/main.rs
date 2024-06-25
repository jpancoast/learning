fn main() {
    let robot = Robot {};
    robot.run();
}

trait Run {
    fn run(&self) {
        println!("I am running");
    }
}

struct Robot {}
impl Run for Robot {}
