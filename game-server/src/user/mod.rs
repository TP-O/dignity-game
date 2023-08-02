use tokio::sync::RwLock;
use std::collections::HashMap;
use std::sync::Arc;

pub mod types;
pub mod routers;
pub mod handle_err;
use crate::user::types::{room::Room, user::User};


#[derive(Debug, Clone)]
pub struct Game {
    pub users: Arc<RwLock<HashMap<usize, User>>>,
    pub connected_users: Arc<RwLock<HashMap<usize, String>>>,
    pub rooms: Arc<RwLock<HashMap<usize, Room>>>,
}

impl Game{
    pub fn new() -> Self{
        Game{
            users: Arc::new(RwLock::new(Self::init())),
            connected_users: Arc::new(RwLock::new(HashMap::new())),
            rooms: Arc::new(RwLock::new(Self::init2())),
        }
    }
    fn init() -> HashMap<usize, User> {
        let file = include_str!("../user/users.json");
        serde_json::from_str(file).expect("can't read users.json")
    }
    fn init2() -> HashMap<usize, Room> {
        let file = include_str!("../user/rooms.json");
        serde_json::from_str(file).expect("can't read rooms.json")
    }
}
