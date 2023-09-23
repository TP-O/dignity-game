
use crate::types::{user::User};
use crate::game_server::Game;
use crate::handle_err::Error;
use serde::{Deserialize, Serialize};
use warp::http::StatusCode;

use super::room::RequestBody;
// use rusty_paseto::{PasetoBuilder, PasetoError, PasetoTimeBackend, PasetoToken};

#[derive(Deserialize, Serialize, Debug, Clone)]
pub struct UserLogin {
    username: String,
    password: String,
}

pub async fn login(
    body: UserLogin,
    game: Game
) -> Result<impl warp::Reply, warp::Rejection> {

    match game.login(body.username,body.password).await {
            Ok(res) => Ok(warp::reply::json(&res)),
            Err(e) => Err(warp::reject::custom(e)),
    }

}



// pub async fn logout(
//     param: String,
//     game: Game
// ) -> Result<impl warp::Reply, warp::Rejection> {
//     let res: Vec<User>= game.users.read().await.values().cloned().collect();
//     if let Some(user) = res.iter().find(|u| u.username == param){
//         match game.connected_users.write().await.remove(&user.id) {
//             Some(_) => (),
//             None => return Err(warp::reject::custom(Error::UserNotFound)),
//         }
//         Ok(warp::reply::with_status("User logout", StatusCode::OK))
//     } else {
//         Err(warp::reject::custom(Error::UserNotFound))
//     }
    
// }
