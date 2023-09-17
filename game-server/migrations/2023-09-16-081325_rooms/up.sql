-- Your SQL goes here
CREATE TABLE IF NOT EXISTS rooms (
  id SERIAL PRIMARY KEY,
  user_id SERIAL REFERENCES users(id)
);
