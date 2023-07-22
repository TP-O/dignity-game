use std::{thread, time::Duration};

mod game;

#[tokio::main]
async fn main() {
    let mut g1 = game::Game::new(1, vec![1, 2, 3, 4, 5]);
    let mut g2 = game::Game::new(2, vec![5, 6, 7, 8, 9]);

    tokio::spawn(async move {
        g1.start().await;
    });

    thread::sleep(Duration::from_secs(6));

    tokio::spawn(async move {
        g2.start().await;
    });

    thread::sleep(Duration::from_secs(99999));
}
