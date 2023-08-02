use serde::{Deserialize, Serialize};
#[derive(Deserialize, Serialize, Debug, Clone)]

pub struct User{
    pub id: usize,
    pub username: String,
    pub room_id: usize 
}