use crate::game::PlayerID;

use super::{Asset, Coin};

pub struct Area {
    owner_id: PlayerID,
    price: Coin,
    buildings: Vec<Building>,
}

impl Area {
    pub fn new(price: Coin) -> Self {
        Self {
            price,
            owner_id: 0,
            buildings: vec![],
        }
    }
}

impl Asset for Area {
    fn category(&self) -> super::AssetCategory {
        super::AssetCategory::Area
    }

    fn price(&self) -> Coin {
        self.price
    }

    fn resale_price(&self) -> Coin {
        (self.price as f32 * 0.8) as Coin
            + self
                .buildings
                .iter()
                .map(|b| b.resale_price())
                .sum::<Coin>()
    }

    fn owner_id(&self) -> PlayerID {
        self.owner_id
    }

    fn set_owner_id(&mut self, player_id: PlayerID) {
        self.owner_id = player_id
    }

    fn is_mortgageable(&self) -> bool {
        self.owner_id != 0
    }
}

pub struct Building {
    owner_id: PlayerID,
    price: Coin,
}

impl Building {
    pub fn new(price: Coin) -> Self {
        Self { price, owner_id: 0 }
    }
}

impl Asset for Building {
    fn category(&self) -> super::AssetCategory {
        super::AssetCategory::Building
    }

    fn price(&self) -> Coin {
        self.price
    }

    fn resale_price(&self) -> Coin {
        (self.price as f32 * 0.8) as Coin
    }

    fn owner_id(&self) -> PlayerID {
        self.owner_id
    }

    fn set_owner_id(&mut self, player_id: PlayerID) {
        self.owner_id = player_id
    }

    fn is_mortgageable(&self) -> bool {
        self.owner_id != 0
    }
}
