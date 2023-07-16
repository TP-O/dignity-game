use std::collections::hash_map::{HashMap, IterMut};

mod paybook;
use self::paybook::{Paybook, PaybookRecord};

use super::{
    asset::{Asset, AssetCategory},
    Coin, GameError, GameResult, ObjectCategory, PlayerID, Ratio, Time,
};

const DEFAULT_INTEREST_RATE: Ratio = 0.1;
const INITIAL_BALANCE: Coin = 10_000_000;

pub struct Bank {
    balance: Coin,
    interest_rate: Ratio,
    paybook: Paybook,
}

impl Bank {
    pub fn new() -> Self {
        Self {
            interest_rate: DEFAULT_INTEREST_RATE,
            balance: INITIAL_BALANCE,
            paybook: HashMap::new(),
        }
    }

    pub fn paybook(&mut self) -> IterMut<'_, PlayerID, Vec<PaybookRecord>> {
        self.paybook.iter_mut()
    }

    pub fn add_balance(&mut self, amount: Coin) {
        self.balance += amount
    }

    pub fn lend(
        &mut self,
        player_id: PlayerID,
        asset: Box<dyn Asset>,
        term: Time,
    ) -> GameResult<Coin> {
        self.lend_with_interest_rate(player_id, asset, term, self.interest_rate)
    }

    pub fn lend_with_interest_rate(
        &mut self,
        player_id: PlayerID,
        asset: Box<dyn Asset>,
        term: Time,
        rate: Ratio,
    ) -> GameResult<Coin> {
        if !asset.is_mortgageable() {
            return Err(GameError::UntradeableAsset {
                category: AssetCategory::Area,
            });
        } else if asset.resale_price() > self.balance {
            return Err(GameError::UnavailableBalance {
                kind: ObjectCategory::Bank,
                balance: self.balance,
            });
        }

        let record = self.paybook.entry(player_id).or_insert(vec![]);
        record.push(PaybookRecord::new(
            player_id,
            asset.resale_price(),
            term,
            if rate >= 0.0 {
                rate
            } else {
                DEFAULT_INTEREST_RATE
            },
        ));
        self.add_balance(asset.resale_price());

        Ok(asset.resale_price())
    }
}
