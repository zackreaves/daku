package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Table interface {
	insert(db_driver string, db_loc string) (sql.Result, error) // Insert into table.
}

type Players struct {
	id uint
	name_first string
}

func (p Players) insert (db_loc string, db_driver string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver, db_loc)
	defer db.Close()
	result, err_exec := db.Exec("INSERT INTO players (names) VALUES ('?')", p.name_first)

	return result, fmt.Errorf("Players Insert Failed: \n%w \n%w\n",err_open,err_exec)
}

type Games struct {
	id uint
	name string
	ties_possible uint8
	tie_breakers uint8
	score_kept uint8
	extensions uint8
}

func (g Games) insert (db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver,db_loc)	
	defer db.Close()
	result, err_exec := db.Exec("INSERT INTO games (name,ties_possible,tie_breakers,score_kept,extensions) VALUES ('?',?,?,?,?);",g.name,g.ties_possible,g.tie_breakers,g.score_kept,g.extensions)


	return result, fmt.Errorf("Games INSERT Failed: \n%w \n%w\n", err_open, err_exec)
}

type Match_data struct {
	game_id uint
	rounds uint
	datetime string
	player_count uint
}

func (m Match_data) insert (db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver,db_loc)
	defer db.Close()
	result, err_exec := db.Exec("INSERT INTO match_data (game_id,rounds,datetime,player_count) VALUES (?,?,'?',?);",m.game_id,m.rounds,m.datetime,m.player_count)

	return result, fmt.Errorf("Match Data INSERT Failed: \n%w \n%w\n",err_open,err_exec)
}

type Player_data struct {
	player_id uint
	match_id uint
	points float64
	win uint8
	ties uint
	round_number uint
}

func (p Player_data) Insert (db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver,db_loc)
	result, err_exec := db.Exec("INSERT INTO player_data (player_id,match_id,points,win,ties,round_number) VALUES (?,?,?,?,?,?);", p.player_id,p.match_id,p.points,p.win,p.ties,p.round_number)

	return result, fmt.Errorf("Player Data INSERT Failed: \n%w \n%w\n",err_open,err_exec)
}

func Insert_to_table (db_driver string, db_loc string, t Table) (sql.Result, error) {
	return t.insert(db_driver, db_loc)
}

func Error_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init_sqlite (db_loc string) {
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
	_, err_match_data := db.Exec(`
		CREATE TABLE IF NOT EXISTS "match_data" (
			"id" INTEGER PRIMARY KEY AUTOINCREMENT,
			"game_id" INTEGER NOT NULL,
			"round_count" INTEGER NOT NULL,
			"player_count" INTEGER NOT NULL,
			"date_time" TEXT,
			FOREIGN KEY("game_id") REFERENCES games("id")
		);
	`)
	Error_check(err_match_data)
	_, err_player_data := db.Exec(`
		CREATE TABLE IF NOT EXISTS "player_data" (
			"match_id" INTEGER NOT NULL,
			"player_id" INTEGER NOT NULL,
		  "win" INTEGER NULL,
		  "score" REAL NULL,
			"tie" BOOLEAN NULL,
			"round_number" INTEGER DEFAULT 1,
		  FOREIGN KEY("match_id") REFERENCES match_data("id"),
		  FOREIGN KEY("player_id") REFERENCES players("id")
		);
	`)
	Error_check(err_player_data)

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
