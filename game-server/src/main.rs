use warp::{Filter};
mod routers;
mod types;
mod game_server;
mod handle_err;
mod models;
mod schema;
mod lib;

#[tokio::main]
async fn main() {

    // let connect = warp::any().map(move || connect);
    let game = game_server::Game::new().await;
    let game_filter = warp::any().map(move || game.clone());


    let test = warp::get().and(warp::path("user")).and(warp::path::end()).map(|| format!("Hello, World!"));

    let login = warp::get().and(warp::path("login")).and(warp::body::json()).and(warp::path::end()).and(game_filter.clone()).and_then(routers::user::login);

    // let logout = warp::get().and(warp::path("logout")).and(warp::path::param::<String>()).and(warp::path::end()).and(connect.clone())
    // .and_then(routers::user::logout);

    // let join_room = warp::post().and(warp::path("join_room")).and(warp::path::end()).and(warp::body::json()).and(data_filter.clone())
    // .and_then(user::routers::room::join_room);

 
    // let routers = 
    warp::serve(login).run(([127, 0, 0, 1], 3030)).await;
}

