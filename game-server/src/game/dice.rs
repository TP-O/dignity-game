use rand::prelude::*;

const DEFAULT_N: i8 = 2;
const DEFAULT_TIMES: i8 = 1;

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

    pub fn roll(&self) -> (Vec<Vec<i8>>, i8) {
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
}
