use std::{thread, time::Duration};
use warp::{http::Method, Filter};

mod game;
mod user;
#[tokio::main]
async fn main() {
    // let mut g1 = game::Game::new(1, vec![1, 2, 3, 4, 5]);
    // let mut g2 = game::Game::new(2, vec![5, 6, 7, 8, 9]);

    // tokio::spawn(async move {
    //     g1.start().await;
    // });

    // thread::sleep(Duration::from_secs(6));

    // tokio::spawn(async move {
    //     g2.start().await;
    // });

    // thread::sleep(Duration::from_secs(99999));
    let game = user::Game::new();
    let data_filter = warp::any().map(move || game.clone());

    let test = warp::get().and(warp::path("user")).and(warp::path::end()).map(|| format!("Hello, World!"));

    let login = warp::get().and(warp::path("login")).and(warp::path::param::<String>()).and(warp::path::end()).and(data_filter.clone())
    .and_then(user::routers::user::login);

    let logout = warp::get().and(warp::path("logout")).and(warp::path::param::<String>()).and(warp::path::end()).and(data_filter.clone())
    .and_then(user::routers::user::logout);

    // let join_room = warp::post().and(warp::path("join_room")).and(warp::path::end()).and(warp::body::json()).and(data_filter.clone())
    // .and_then(user::routers::room::join_room);

    let routers = test.or(login).or(logout);
    warp::serve(routers).run(([127, 0, 0, 1], 3030)).await;
}
