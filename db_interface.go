package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func error_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init (db_loc string) {
	db, err_open := sql.Open("sqlite3", db_loc)

	error_check(err_open)

	defer db.Close()

	_, err_players := db.Exec(`
		CREATE TABLE IF NOT EXISTS "players" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name_first" VARCHAR(80)
		);
	`)
	error_check(err_players)
	_, err_games := db.Exec(`
		CREATE TABLE IF NOT EXISTS "games" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name" VARCHAR(80),
			"ties_possible" BOOLEAN,
			"tie_breakers" BOOLEAN,
			"score_kept" BOOLEAN,
			"round_extensions" BOOLEAN
		);
	`)
	error_check(err_games)
	// FIXME: Add extra columns to layout.
	_, err_round_data := db.Exec(`
		CREATE TABLE IF NOT EXISTS "round_data" (
		  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"game_id" INTEGER NOT NULL,
		  "round_count" INTEGER NOT NULL,
		  "player_count" INTEGER NOT NULL,
		  "ties" INTEGER NULL,
		  "date_time" TEXT,
		  FOREIGN KEY("game_id") REFERENCES games("id")
		);
	`)
	error_check(err_round_data)
	_, err_player_data := db.Exec(`
		CREATE TABLE IF NOT EXISTS "player_data" (
			"round_id" INTEGER NOT NULL,
			"player_id" INTEGER NOT NULL,
		  "win" INTEGER NOT NULL,
		  "score" REAL NULL,
			"tie" BOOLEAN NULL,
			"round_number" INTEGER NULL,
		  FOREIGN KEY("round_id") REFERENCES round_data("id"),
		  FOREIGN KEY("player_id") REFERENCES players("id")
		);
	`)
	error_check(err_player_data)


	return
}

func Exec(db_loc string, query string) sql.Result {
	db, err_open := sql.Open("sqlite3",db_loc)

	error_check(err_open)

	defer db.Close()

	result, err_query := db.Exec(query)

	error_check(err_query)

	return result
}
