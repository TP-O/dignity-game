use super::{asset::Asset, Coin, PlayerID};

const INITIAL_BALANCE: Coin = 1_000_000;

pub struct Player {
    id: PlayerID,
    balance: Coin,
}

impl Player {
    pub fn new(id: PlayerID) -> Self {
        Self {
            id,
            balance: INITIAL_BALANCE,
        }
    }

    pub fn id(&self) -> PlayerID {
        self.id
    }

    pub fn can_buy(&self, asset: Box<dyn Asset>) -> bool {
        self.balance >= asset.price() as i64
    }

    pub fn add_balance(&mut self, amount: Coin) -> bool {
        self.balance += amount;
        self.balance > 0
    }
}
