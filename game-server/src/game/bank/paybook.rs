use std::collections::HashMap;

use crate::game::{Coin, GameError, GameResult, PlayerID, Ratio, Time};

pub type Paybook = HashMap<PlayerID, Vec<PaybookRecord>>;

pub struct PaybookRecord {
    player_id: PlayerID,
    initial_debt: Coin,
    current_debt: Coin,
    term: Time,
    interval: i64,
    interval_counter: i64,
    interest_rate: Ratio,
    is_late: bool,
}

impl PaybookRecord {
    pub fn new(player_id: PlayerID, amount: Coin, term: Time, rate: Ratio) -> Self {
        Self {
            player_id,
            initial_debt: amount,
            current_debt: amount,
            interval: match term {
                Time::Round(_) => 1,
                Time::Turn(_) => 4,  // 4 turns ~ 1 round
                Time::Step(_) => 48, // 12 steps ~ 1 turn
            },
            term,
            interval_counter: 0,
            interest_rate: rate,
            is_late: false,
        }
    }

    pub fn is_done(&self) -> bool {
        self.current_debt == 0
    }

    pub fn payment_due(&self) -> Coin {
        (self.initial_debt / self.interval)
            + (self.current_debt as f32 * self.interest_rate) as Coin
    }

    pub fn pay(&mut self) -> GameResult<Coin> {
        if self.is_done() {
            return Err(GameError::Denied("Debt has been paid in full".to_owned()));
        } else if !self.is_due() {
            return Err(GameError::Denied("Please wait until due".to_owned()));
        }

        self.current_debt -= self.payment_due();
        self.interval_counter = 0;

        Ok(self.current_debt)
    }

    pub fn delay(&mut self) -> GameResult<Ratio> {
        if self.is_late {
            return Err(GameError::Denied("Payment deferral is denied".to_owned()));
        } else if !self.is_due() {
            return Err(GameError::Denied("Please wait until due".to_owned()));
        }

        self.is_late = true;
        self.interest_rate *= 2.0;

        Ok(self.interest_rate)
    }

    fn is_due(&self) -> bool {
        return self.interval == self.interval_counter;
    }
}
