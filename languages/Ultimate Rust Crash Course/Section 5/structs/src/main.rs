fn main() {
    let fox = RedFox::new();
    println!("{}, {}", fox.life, fox.enemy);

    print_noise(fox);
}

struct RedFox {
    enemy: bool,
    life: i32,
}

impl RedFox {
    fn new() -> Self {
        Self {
            enemy: true,
            life: 70,
        }
    }
}

trait Noisy {
    fn get_noise(&self) -> &str;
}

impl Noisy for RedFox {
    fn get_noise(&self) -> &str {
        "Bark"
    }
}

/*
 * Why use traits?
 *
 * So we can write a generic function that can accept anything that has the"Noisy" trait, like so
 */
fn print_noise<T: Noisy>(item: T) {
    println!("{}", item.get_noise());
}
