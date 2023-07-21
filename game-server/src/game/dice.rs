use rand::prelude::*;
use tokio::time;
const DEFAULT_N: i8 = 2;
const DEFAULT_TIMES: i8 = 1;

use tokio::time::{Duration,sleep};
use tokio::io::{self, AsyncBufReadExt};
use crossterm::event::{
    self as crosstermEvent, read, DisableMouseCapture, EnableMouseCapture, KeyCode, KeyEvent, KeyModifiers,
};
pub struct Dice {
    n: i8,
    times: i8,
}

impl Default for Dice {
    fn default() -> Self {
        Self {
            n: DEFAULT_N,
            times: DEFAULT_TIMES,
        }
    }
}

impl Dice {
    pub fn reset(&mut self) -> &mut Self {
        self.n = DEFAULT_N;
        self.times = DEFAULT_TIMES;
        self
    }

    pub fn set_n(&mut self, n: i8) -> &mut Self {
        self.n = n;
        self
    }

    pub fn set_times(&mut self, times: i8) -> &mut Self {
        self.times = times;
        self
    }

    pub async fn roll(&self) -> (Vec<Vec<i8>>, i8) {
        let mut result = (Vec::<Vec<i8>>::new(), 0);
        let mut rng = rand::thread_rng();

        for _ in 0..self.times {
            let mut faces = vec![];
            for _ in 0..self.n {
                faces.push(rng.gen_range(1..7));
                result.1 += faces.last().unwrap();
            }

            result.0.push(faces);
        }
        result
    }
    async fn read_input(
        reader: &mut io::BufReader<io::Stdin>,
    ) -> Result<Option<String>, Box<dyn std::error::Error>> {
        loop {
            if crosstermEvent::poll(Duration::from_millis(100))? {
                if let crosstermEvent::Event::Key(KeyEvent {
                    code, modifiers, ..
                }) = read()?
                {
                    if code != KeyCode::Char('\n') || modifiers != KeyModifiers::NONE {
                        let mut input = String::new();
                        reader.read_line(&mut input).await?;
                        return Ok(Some(input.trim().to_string()));
                    } else {
                        return Ok(None);
                    }
                }
            }
        }
    }
    pub async fn run(&self) -> (Vec<Vec<i8>>, i8){
        let _mouse_capture_guard = EnableMouseCapture;
        let stdin = io::stdin();
        let mut reader = io::BufReader::new(stdin);
        println!("Guess roll!");
        loop {
            let mut guess: String = String::new();
            println!("Please roll within the 15 seconds:");
            let timer = sleep(Duration::from_secs(15));
            tokio::select! {
                _ = timer => {
                    println!("Time's up! You did not roll.");
                    return (vec![], 0);
                }
                result = Self::read_input(&mut reader) => {
                    match result {
                        Ok(Some(guess)) => {
                            let (dices, steps) = Self::roll(self).await;
                            return (dices, steps);
                        },
                        Ok(None) => println!("None"),
                        Err(e) => eprintln!("Error: {}", e),
                    }
                }
            }
        }
    }
}
