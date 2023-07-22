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
use tokio::sync::watch::{self, Receiver, Sender};
use tokio::time::{sleep, Duration};

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
    id: i64,
    dice: Dice,
    map: Map,
    event: Event,
    lowest_balance: Coin,
    players: Vec<Player>,
    bank: Bank,
    turn_channel: (Sender<i64>, Receiver<i64>),
}

impl Game {
    pub fn new(id: i64, player_ids: Vec<PlayerID>) -> Self {
        Self {
            id,
            lowest_balance: 0,
            dice: Dice::default(),
            players: player_ids.into_iter().map(|id| Player::new(id)).collect(),
            bank: Bank::new(),
            map: Map::new(),
            event: Event::new(),
            turn_channel: watch::channel(0),
        }
    }

    pub async fn start(&mut self) {
        loop {
            for player in self.players.iter_mut() {
                let timer = sleep(Duration::from_secs(5));

                tokio::select! {
                    _ = timer => {
                        let (_, steps) = self.dice.roll();
                        println!("[G{}]time's up!", self.id);
                        println!("[G{}] machine roll for player {}: {} steps", self.id, player.id(), steps);
                    }
                    _ = self.turn_channel.1.changed() => {
                        let (_, steps) = self.dice.roll();
                        println!("[G{}]player {} moves {} steps", self.id, player.id(), steps);
                    }
                }
            }
        }
    }
}
