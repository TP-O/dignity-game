mod game;

fn main() {
    let mut g = game::Game::new(vec![1, 2, 3, 4, 5]);
    g.start();
}
