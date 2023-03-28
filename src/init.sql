CREATE TABLE IF NOT EXISTS "players" (
  "id" INTEGER PRIMARY KEY NOT NULL,
  "name_first" VARCHAR(80)
);

CREATE TABLE IF NOT EXISTS "games" (
  "id" INTEGER PRIMARY KEY NOT NULL,
  "name" VARCHAR(80),
  "ties_possible" BOOLEAN,
  "score_kept" BOOLEAN
);

CREATE TABLE IF NOT EXISTS "round_data" (
  "id" INTEGER PRIMARY KEY NOT NULL,
  FOREIGN KEY("game_id") REFERENCES games("id"),
  "round_count" INTEGER,
  "player_count" INTEGER,
  "ties" INTEGER FALSE,
  "date_time" DATETIME
);

CREATE TABLE IF NOT EXISTS "player_data" (
  FOREIGN KEY("round_id") REFERENCES round_data("id"),
  FOREIGN KEY("player_id") REFERENCES players("id"),
  "wins" INTEGER NOT NULL,
  "score" INTEGER NULL
);