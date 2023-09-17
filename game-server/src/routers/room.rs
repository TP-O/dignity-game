use crate::types::room::Room;
use crate::game_server::Game;
use crate::handle_err::Error;

use serde::Deserialize;

#[derive(Deserialize)]
pub struct RequestBody {
    user_id: usize,
    room_id: usize,
}

// pub async fn join_room(body: RequestBody, game: Game) -> Result<impl warp::Reply, warp::Rejection> {
//     let res: Vec<Room> = game.rooms.read().await.values().cloned().collect();
//     if let Some(room) = res.iter().find(|r| r.id == body.room_id) {
//         if let Some(player_index) = room.players.iter().position(|&p| p == 0) {
//             let mut room = room.clone();
//             room.players[player_index] = body.user_id;
//             game.rooms.write().await.get_mut(&body.room_id);
//             Ok(warp::reply::json(&room))
//         } else {
//             Err(warp::reject::custom(Error::RoomFull))
//         }
//     } else {
//         Err(warp::reject::custom(Error::RoomNotFound))
//     }
// }
// pub async fn check_connect(user: User, game:Game) -> Result<impl warp::Reply, warp::Rejection>{
//     let res: Vec<Room>= game.rooms.read().await.values().cloned().collect();
//     if user.room_id != 0 {
//         join_room(user.id, user.room_id, game);

//     }
//     else {
//         Ok(warp::reply::json(&user))
//     }
// }
