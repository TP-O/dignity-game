
use crate::user::types::{user::User};
use crate::user::Game;
use crate::user::handle_err::Error;
use warp::http::StatusCode;
// use rusty_paseto::{PasetoBuilder, PasetoError, PasetoTimeBackend, PasetoToken};



pub async fn login(
    param: String,
    game: Game
) -> Result<impl warp::Reply, warp::Rejection> {
    let res: Vec<User>= game.users.read().await.values().cloned().collect();
    let connected_user = game.connected_users.read().await;
    if let Some(user) = res.iter().find(|u| u.username == param){
        match connected_user.get(&user.id){
            Some(_) =>  Err(warp::reject::custom(Error::AccountUsedInGame)),
            None => {
                drop(connected_user);
                game.connected_users.write().await.insert(user.id.clone(),"connected".to_string());
                Ok(warp::reply::json(&user))
            }
        }
    } else {
        Err(warp::reject::custom(Error::UserNotFound))
    }
}



pub async fn logout(
    param: String,
    game: Game
) -> Result<impl warp::Reply, warp::Rejection> {
    let res: Vec<User>= game.users.read().await.values().cloned().collect();
    if let Some(user) = res.iter().find(|u| u.username == param){
        match game.connected_users.write().await.remove(&user.id) {
            Some(_) => (),
            None => return Err(warp::reject::custom(Error::UserNotFound)),
        }
        Ok(warp::reply::with_status("User logout", StatusCode::OK))
    } else {
        Err(warp::reject::custom(Error::UserNotFound))
    }
    
}
