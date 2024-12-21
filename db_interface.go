package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Table interface {
	Insert(db_driver string, db_loc string) (sql.Result, error) // Insert into table.
	Populate_from_args(args []string, format []string)
}

type Players struct {
	id uint
	name_first string
}

func (p *Players) Populate_from_args(args []string, format []string) {
	for i := 0; i < len(args); i++ {
		switch format[i] {
		case "id":
			id,_ := strconv.ParseUint(args[i],10,64)	
			p.id = uint(id)
		case "name_first":
			p.name_first = args[i]
		}
	}
}

func (p Players) Insert(db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver, db_loc)
	defer db.Close()
	result, err_exec := db.Exec("INSERT INTO players (name_first) VALUES ($1)", p.name_first)

	return result, fmt.Errorf("Players INSERT Failed: \nOn Open: %w\nOn INSERT: %w\nAttempted Value: %s\n",err_open,err_exec, p.name_first)
}

type Games struct {
	id uint
	name string
	ties_possible bool
	tie_breakers bool
	score_kept bool
	extensions bool
}

func (g *Games) Populate_from_args (args []string, format []string) {
	for i := 0; i < len(args); i++ {
		switch format[i] {
		case "name":
			g.name = args[i]
		case "ties_possible":
			g.ties_possible,_ = strconv.ParseBool(args[i])
		case "score_kept":
			g.score_kept,_ = strconv.ParseBool(args[i])
		case "tie_breakers":
			g.tie_breakers,_ = strconv.ParseBool(args[i])
		case "extensions":
			g.extensions,_ = strconv.ParseBool(args[i])
		case "id":
			id,_ := strconv.ParseUint(args[i],10,64)
			g.id = uint(id)
		}
	}
}

func (g Games) Insert (db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver,db_loc)	
	defer db.Close()
	result, err_exec := db.Exec("INSERT INTO games (name,ties_possible,tie_breakers,score_kept,extensions) VALUES ($1,$2,$3,$4,$5);",g.name,g.ties_possible,g.tie_breakers,g.score_kept,g.extensions)


	return result, fmt.Errorf("Games INSERT Failed: \n%w \n%w\n", err_open, err_exec)
}

type Match_data struct {
	id uint
	game_id uint
	rounds uint
	datetime string
	player_count uint
}

func (m *Match_data) Populate_from_args (args []string, format []string) {
	for i := 0; i < len(args); i++ {
		switch format[i] {
		case "id":
			id,_ := strconv.ParseUint(args[i],10,64)
			m.id = uint(id)
		case "game_id":
			id,_ := strconv.ParseUint(args[i],10,64)
			m.game_id = uint(id)
		case "rounds":
			rounds,_ := strconv.ParseUint(args[i],10,64)
			m.rounds = uint(rounds)
		case "datetime":
			m.datetime = args[i]
		case "player_count":
			count,_ := strconv.ParseUint(args[i],10,64)
			m.player_count = uint(count)
		}
	}
}

func (m Match_data) Insert (db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver,db_loc)
	defer db.Close()
	result, err_exec := db.Exec("INSERT INTO match_data (game_id,rounds,datetime,player_count) VALUES ($1,$2,$3,$4);",m.game_id,m.rounds,m.datetime,m.player_count)

	return result, fmt.Errorf("Match Data INSERT Failed: \n%w \n%w\n",err_open,err_exec)
}

type Player_data struct {
	player_id uint
	match_id uint
	points float64
	win bool
	ties uint
	round_number uint
}

func (p *Player_data) Populate_from_args (args []string, format []string) {
	for i := 0; i < len(args); i++ {
		switch format[i] {
		case "player_id":
			vuint,_ := strconv.ParseUint(args[i],10,64)
			p.player_id = uint(vuint)
		case "match_id":
			vuint,_ := strconv.ParseUint(args[i],10,64)
			p.match_id = uint(vuint)
		case "points":
			vfloat,_ := strconv.ParseFloat(args[i],64)
			p.points = vfloat
		case "round_number":
			vuint,_ := strconv.ParseUint(args[i],10,64)
			p.round_number = uint(vuint)
		case "win":
			vbool,_ := strconv.ParseBool(args[i])
			p.win = vbool
		}
	}
}

func (p Player_data) Insert (db_driver string, db_loc string) (sql.Result, error) {
	db, err_open := sql.Open(db_driver,db_loc)
	result, err_exec := db.Exec("INSERT INTO player_data (player_id,match_id,points,win,ties,round_number) VALUES ($1,$2,$3,$4,$5,$6);", p.player_id,p.match_id,p.points,p.win,p.ties,p.round_number)

	return result, fmt.Errorf("Player Data INSERT Failed: \n%w \n%w\n",err_open,err_exec)
}

func Populate_from_arguments (args []string, format []string, t Table) {
	t.Populate_from_args(args,format)
}

func Insert_from_table (db_driver string, db_loc string, t Table) (sql.Result, error) {
	return t.Insert(db_driver, db_loc)
}

func Error_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init (db_driver string, db_loc string) {
	db, err_open := sql.Open(db_driver,"file:" + db_loc + "?_foreign_keys=true")

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

	if db_driver == "sqlite3" {
		_, err_match_data := db.Exec(`
			CREATE TABLE IF NOT EXISTS "match_data" (
				"id" INTEGER PRIMARY KEY AUTOINCREMENT,
				"game_id" INTEGER NOT NULL,
				"round_count" INTEGER NOT NULL,
				"player_count" INTEGER NOT NULL,
				"date_time" DATETIME DEFAULT datetime('now'),
				FOREIGN KEY("game_id") REFERENCES games("id")
			);
		`)
		Error_check(err_match_data)
	} else {
		_, err_match_data := db.Exec(`
			CREATE TABLE IF NOT EXISTS "match_data" (
				"id" INTEGER PRIMARY KEY AUTOINCREMENT,
				"game_id" INTEGER NOT NULL,
				"round_count" INTEGER NOT NULL,
				"player_count" INTEGER NOT NULL,
				"date_time" DATETIME DEFAULT NOW(),
				FOREIGN KEY("game_id") REFERENCES games("id")
			);
		`)
		Error_check(err_match_data)
	}
	_, err_player_data := db.Exec(`
		CREATE TABLE IF NOT EXISTS "player_data" (
			"match_id" INTEGER NOT NULL,
			"player_id" INTEGER NOT NULL,
			"win" BOOLEAN NULL,
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
