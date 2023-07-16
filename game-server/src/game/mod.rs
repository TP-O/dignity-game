mod dice;
use dice::Dice;

mod player;
use player::Player;

mod map;
use map::Map;

mod event;
use event::Event;

mod asset;

mod bank;
use bank::Bank;

pub type Coin = i64;

pub type PlayerID = u64;

pub type Ratio = f32;

pub enum Time {
    Turn(u64),
    Step(u64),
    Round(u64),
}

pub enum GameError {
    UntradeableAsset { category: asset::AssetCategory },
    UnavailableBalance { kind: ObjectCategory, balance: Coin },
    Denied(String),
}

pub enum ObjectCategory {
    Bank,
    Player(u64),
}

type GameResult<T> = Result<T, GameError>;

pub struct Game {
    dice: Dice,
    map: Map,
    event: Event,
    lowest_balance: Coin,
    players: Vec<Player>,
    bank: Bank,
}

impl Game {
    pub fn new(player_ids: Vec<PlayerID>) -> Self {
        Self {
            lowest_balance: 0,
            dice: Dice::default(),
            players: player_ids.into_iter().map(|id| Player::new(id)).collect(),
            bank: Bank::new(),
            map: Map::new(),
            event: Event::new(),
        }
    }

    pub fn start(&mut self) {
        loop {
            for player in self.players.iter_mut() {
                let (_, steps) = self.dice.roll();
                println!("Player {} moves {} steps", player.id(), steps);
            }
        }
    }
}
