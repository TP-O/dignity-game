mod game;


#[tokio::main]
async fn main() {
    let mut g = game::Game::new(vec![1, 2, 3, 4, 5]);
    let mut g1 = game::Game::new(vec![5, 6,7, 8, 9]);
    g.start().await;
    // let task2_handle = tokio::task::spawn(async move {g1.start().await});

    // Đợi cả hai nhiệm vụ hoàn thành
}
