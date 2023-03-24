--CREATE DATABASE daku_scorekeeper

CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  namef VARCHAR(80),
);

CREATE TABLE games (
  id SERIAL PRIMARY KEY,
  name VARCHAR(80),
  ties BOOLEAN,
);

CREATE TABLE round_data (
  id SERIAL PRIMARY KEY,
  player_id INTEGER REFERENCES players(id),
  round_count INTEGER, -- FIXME: Add some automated logic that sums the rounds based on player data.
  ties INTEGER NULL,
  player_count INTEGER,
  -- FIXME: Add date and time.
);

CREATE TABLE player_data (
  round_id INTEGER REFERENCES round_data(id),
  player_id INTEGER REFERENCES players(id),
  wins INTEGER,
  score INTEGER,
)
