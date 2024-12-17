package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Players struct {
	name_first string
}

type Games struct {
	name string
	ties_possible uint8
	tie_breakers uint8
	score_kept uint8
	extensions uint8
}

type Round_data struct {
	rounds uint
	ties uint
	datetime string
	game string
	player_count uint
}

type Player_data struct {
	name string
	points float64
	win uint8 // should be bool
	ties uint
	round_number uint
}

func Error_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init (db_loc string) {
	db, err_open := sql.Open("sqlite3","file:" + db_loc + "?_foreign_keys=true")

	Error_check(err_open)

	defer db.Close()

	_, err_players := db.Exec(`
		CREATE TABLE IF NOT EXISTS "players" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"name_first" VARCHAR(80)
		);
	`)
	Error_check(err_players)
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
	Error_check(err_games)
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
	Error_check(err_round_data)
	_, err_player_data := db.Exec(`
		CREATE TABLE IF NOT EXISTS "player_data" (
			"round_id" INTEGER NOT NULL,
			"player_id" INTEGER NOT NULL,
		  "win" INTEGER NOT NULL,
		  "score" REAL NULL,
			"tie" BOOLEAN NULL,
			"round_number" INTEGER NOT NULL,
		  FOREIGN KEY("round_id") REFERENCES round_data("id"),
		  FOREIGN KEY("player_id") REFERENCES players("id")
		);
	`)
	Error_check(err_player_data)


	return
}

func Exec(db_loc string, query string) sql.Result {
	db, err_open := sql.Open("sqlite3","file:" + db_loc + "?_foreign_keys=true")

	Error_check(err_open)

	defer db.Close()

	result, err_query := db.Exec(query)

	Error_check(err_query)

	return result
}

func Query(db_loc string, query string) *sql.Rows {
	db, err_open := sql.Open("sqlite3","file:" + db_loc + "?_foreign_keys=true")

	Error_check(err_open)

	defer db.Close()

	result, err_query := db.Query(query)

	Error_check(err_query)

	return result
}
