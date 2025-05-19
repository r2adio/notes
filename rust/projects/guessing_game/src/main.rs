use std::io;

use rand::Rng;

fn main() {
    println!("The Guessing Game");

    let secret_num = rand::thread_rng().gen_range(1..=100);
    println!("the secret number is {secret_num}");

    println!("enter a value: ");
    let mut guess = String::new();
    io::stdin()
        .read_line(&mut guess)
        .expect("failed to read line");

    println!("you guessed {}", guess);
}
