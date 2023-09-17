use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize, Debug, Clone)]
pub struct Room{
    pub id : usize,
    pub players: [usize;5]
}