use super::{Coin, PlayerID};

pub mod area;

pub enum AssetCategory {
    Area,
    Building,
}

pub trait Asset {
    fn category(&self) -> AssetCategory;
    fn price(&self) -> Coin;
    fn resale_price(&self) -> Coin;
    fn owner_id(&self) -> PlayerID;
    fn set_owner_id(&mut self, player_id: PlayerID);
    fn is_mortgageable(&self) -> bool;
}
