package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init (db_loc string) {
	db, err := sql.Open("sqlite3", db_loc)

	defer db.Close()

	db.Exec(`
		CREATE TABLE IF NOT EXISTS "players" (
			"id" INTEGER PRIMARY KEY NOT NULL,
			"name_first" VARCHAR(80)
		);
	`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS "games" (
			"id" INTEGER PRIMARY KEY NOT NULL,
			"name" VARCHAR(80),
			"ties_possible" BOOLEAN,
			"tie_breakers" BOOLEAN,
			"score_kept" BOOLEAN,
			"round_extensions" BOOLEAN
		);
	`)
	// FIXME: Add extra columns to layout.
	db.Exec(`
		CREATE TABLE IF NOT EXISTS "round_data" (
		  "id" INTEGER PRIMARY KEY NOT NULL,
		  FOREIGN KEY("game_id") REFERENCES games("id"),
		  "round_count" INTEGER,
		  "player_count" INTEGER,
		  "ties" INTEGER NULL,
		  "date_time" TEXT
		);
	`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS "player_data" (
		  "win" INTEGER NOT NULL,
		  "score" REAL NULL,
			"tie" BOOLEAN NULL,
			"round_number" INTEGER NULL,
		  FOREIGN KEY("round_id") REFERENCES round_data("id"),
		  FOREIGN KEY("player_id") REFERENCES players("id")
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	return
}
