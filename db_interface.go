package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type Table interface {
	Insert(db_driver string, db_loc string) (error) // Insert into table.
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

func (p Players) Insert (db_driver string, db_loc string) (error) {
	db, err_open := sql.Open(db_driver, db_loc)
	defer db.Close()

	if err_open != nil {
		return err_open
	}

	_, err_exec := db.Exec("INSERT INTO players (name_first) VALUES ($1)", p.name_first)

	return err_exec
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
		case "round_extensions":
			g.extensions,_ = strconv.ParseBool(args[i])
		case "id":
			id,_ := strconv.ParseUint(args[i],10,64)
			g.id = uint(id)
		}
	}
}

func (g Games) Insert (db_driver string, db_loc string) (error) {
	db, err_open := sql.Open(db_driver,db_loc)	
	defer db.Close()
	if err_open != nil {
		return err_open
	}

	_, err_exec := db.Exec("INSERT INTO games (name,ties_possible,tie_breakers,score_kept,round_extensions) VALUES ($1,$2,$3,$4,$5);",g.name,g.ties_possible,g.tie_breakers,g.score_kept,g.extensions)

	return err_exec
}

type Match_data struct {
	id uint
	game_id uint
	round_count uint
	date_time string
	player_count uint
	relative_id bool
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
		case "round_count":
			rounds,_ := strconv.ParseUint(args[i],10,64)
			m.round_count = uint(rounds)
		case "date_time":
			m.date_time = args[i]
		case "player_count":
			count,_ := strconv.ParseUint(args[i],10,64)
			m.player_count = uint(count)
		case "relative_id":
			vbool,_ := strconv.ParseBool(args[i])
			m.relative_id = vbool
		}
	}
}

func (m Match_data) Insert (db_driver string, db_loc string) (error) {
	if m.relative_id {
		return fmt.Errorf("INSERT INTO match_data FAILED: only absolute id is allowed.")
	}

	db, err_open := sql.Open(db_driver,db_loc)
	defer db.Close()
	if err_open != nil {
		return err_open
	}

	_, err_exec := db.Exec("INSERT INTO match_data (game_id,round_count,date_time,player_count) VALUES ($1,$2,$3,$4);",m.game_id,m.round_count,m.date_time,m.player_count)

	return err_exec
}

type Player_data struct {
	player_id uint
	match_id uint
	score float64
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
		case "score":
			vfloat,_ := strconv.ParseFloat(args[i],64)
			p.score = vfloat
		case "round_number":
			vuint,_ := strconv.ParseUint(args[i],10,64)
			p.round_number = uint(vuint)
		case "win":
			vbool,_ := strconv.ParseBool(args[i])
			p.win = vbool
		case "ties":
			vuint,_ := strconv.ParseUint(args[i],10,64)
			p.ties = uint(vuint)
		}
	}
}

func (p Player_data) Insert (db_driver string, db_loc string) (error) {

	db, err_open := sql.Open(db_driver,db_loc)
	defer db.Close()

	if err_open != nil {
		return err_open
	} 

	_, err_exec := db.Exec("INSERT INTO player_data (player_id,match_id,score,win,ties,round_number) VALUES ($1,$2,,$3,$4,$5,$6);", p.player_id,p.match_id,p.score,p.win,p.ties,p.round_number)

	return err_exec
}

func Populate_from_arguments (args []string, format []string, t Table) {
	t.Populate_from_args(args,format)
}

func Insert_from_table (db_driver string, db_loc string, t Table) (error) {
	return t.Insert(db_driver, db_loc)
}

func Csv_insert (csv_file string, table_type string) { // Might rip this element out of this function later, since I don't know if it's going to be used again.
	var t Table
	csv_arr, rows := Import_from_csv(csv_file)
	format := csv_arr[0]
	csv_args := csv_arr[1:]
	switch table_type {
	case "players":
		t = &Players{}
	case "games":
		t = &Games{}
	case "match_data":
		t = &Match_data{}
	case "player_data":
		t = &Player_data{}
	}
	for i := 0; i < rows-1 ; i++ {
		Populate_from_arguments(csv_args[i], format, t)
		err := Insert_from_table(config.db_driver,config.db_address,t)
		Error_check(err)
	}
}

func Match_populate (matches_csv string, players_csv string) ([]Match_data, []Player_data) {
	match_arr, match_rows := Import_from_csv(matches_csv)
	player_arr, player_rows := Import_from_csv(players_csv)
	matches := make([]Match_data,match_rows)
	players := make([]Player_data,player_rows)

	match_format := match_arr[0]
	match_args := match_arr[1:]

	for i := 0; i < match_rows-1; i++ {
		matches[i].Populate_from_args(match_args[i],match_format)
	}

	player_format := player_arr[0]
	player_args := player_arr[1:]

	for j := 0; j < player_rows-1; j++ {
		players[j].Populate_from_args(player_args[j],player_format)
	}

	return matches, players
}

func Match_sort_insert (config Settings, matches []Match_data, players []Player_data) {

		db, err := sql.Open(config.db_driver,config.db_address)
		Error_check(err)
		defer db.Close()

		tx, err := db.Begin()
		Error_check(err)
		defer tx.Rollback()

		match_stmt, err := tx.Prepare("INSERT INTO match_data (game_id,round_count,date_time,player_count) VALUES ($1,$2,$3,$4);")
		Error_check(err)

		player_stmt, err := tx.Prepare("INSERT INTO player_data (match_id,player_id,win,score,ties,round_number) VALUES ((SELECT MAX (id) FROM match_data),$1,$2,$3,$4,$5);")
		Error_check(err)

		defer match_stmt.Close()
		defer player_stmt.Close()

		for i := 0; i < len(matches)-1; i++ {
			_, err := match_stmt.Exec(matches[i].game_id, matches[i].round_count, matches[i].date_time, matches[i].player_count)
			Error_check(err)
			for j := 0; j < len(players)-1; j++ {
				if players[j].match_id == matches[i].id {
					_,err = player_stmt.Exec(players[j].player_id, players[j].win, players[j].score, players[j].ties, players[j].round_number)
					Error_check(err)
				}
			}
		}

		Error_check(tx.Commit())
}

func Error_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init (config Settings) {
	fmt.Println(config.db_driver)
	fmt.Println(config.db_address)
	switch config.db_driver {
	case "sqlite3":
	db, err_open := sql.Open(config.db_driver,"file:" + config.db_address + "?_foreign_keys=true")

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
				"date_time" DATETIME DEFAULT (datetime('now','localtime')),
				FOREIGN KEY("game_id") REFERENCES games("id")
			);
		`)
		Error_check(err_match_data)
		_, err_player_data := db.Exec(`
			CREATE TABLE IF NOT EXISTS "player_data" (
				"match_id" INTEGER DEFAULT (SELECT last_insert_rowid FROM match_data),
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
	case "postgres":
		db, err_open := sql.Open(config.db_driver,config.db_address)

		Error_check(err_open)

		defer db.Close()

		_, err_exec := db.Exec(`
			CREATE TABLE IF NOT EXISTS "players" (
				"id" SERIAL PRIMARY KEY,
				"name_first" VARCHAR(80)
			);
			CREATE TABLE IF NOT EXISTS "games" (
				"id" SERIAL PRIMARY KEY,
				"name" VARCHAR(80),
				"ties_possible" BOOLEAN,
				"tie_breakers" BOOLEAN,
				"score_kept" BOOLEAN,
				"round_extensions" BOOLEAN
			);
			CREATE TABLE IF NOT EXISTS "match_data" (
				"id" SERIAL PRIMARY KEY, 
				"game_id" INTEGER REFERENCES games(id),
				"round_count" INTEGER NOT NULL,
				"player_count" INTEGER NOT NULL,
				"date_time" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			);
			CREATE TABLE IF NOT EXISTS "player_data" (
				"match_id" INTEGER REFERENCES match_data(id) ON DELETE CASCADE,
				"player_id" INTEGER REFERENCES players(id),
				"win" BOOLEAN NULL,
				"score" REAL NULL,
				"ties" REAL NULL,
				"round_number" INTEGER DEFAULT 1
			);
		`)
		Error_check(err_exec)
	default:
		fmt.Println("Supported databases include postgres and sqlite3.\n You can set the database using the DAKU_SQL_SERVICE environment variable.")
	}
}

func Query (config Settings, query string) *sql.Rows {

	var ( 
		db *sql.DB
		err_open error	
	)

	switch config.db_driver {
	case "sqlite3":
		db, err_open = sql.Open("sqlite3","file:" + config.db_address + "?_foreign_keys=true")
	case "postgres":
		db, err_open = sql.Open("postgres",config.db_address)
	}

	Error_check(err_open)

	defer db.Close()

	result, err_query := db.Query(query)

	Error_check(err_query)

	return result
}

func Query_name (config Settings) ([]Players, []string) {
	query := "SELECT * FROM players;"
	result := Query(config, query)
	var (
		player Players
		players []Players
	)

	defer result.Close()

	columns,_ := result.Columns()

	for result.Next() {
		result.Scan(&player.id, &player.name_first)
		players = append(players, player)
	}

	return players, columns
}

func Query_games (config Settings) ([]Games, []string) {
	query := "SELECT * FROM games;"
	result := Query(config, query)
	var (
		game Games
		games []Games
	)

	defer result.Close()

	columns,_ := result.Columns()

	for result.Next() {
		result.Scan(&game.id,&game.name,&game.ties_possible,&game.tie_breakers,&game.score_kept,&game.extensions)
		games = append(games,game)
	}

	return games, columns
}

type Collated_player_stats struct {
	name string
	win_rate float64
	avg_score float64
}

func Query_players_all (config Settings) []Players {
	var player Players
	var players []Players

	db, err := sql.Open(config.db_driver,config.db_address)
	defer db.Close()
	Error_check(err)

	query_result, err := db.Query("SELECT * FROM players;")
	defer query_result.Close()
	Error_check(err)

	for query_result.Next() {
		err = query_result.Scan(&player.id,&player.name_first)
		Error_check(err)
		players = append(players, player)
	}

	return players
}

func Query_games_all (config Settings) []Games {
	var (
		game Games
		games []Games
		query_result *sql.Rows
	)

	db, err := sql.Open(config.db_driver,config.db_address)
	defer db.Close()
	Error_check(err)

	query_result, err = db.Query("SELECT * FROM games;")
	defer query_result.Close()
	Error_check(err)

	for query_result.Next() {
		err = query_result.Scan(&game.id,&game.name,&game.ties_possible,&game.tie_breakers,&game.score_kept,&game.extensions)
		Error_check(err)
		games = append(games, game)
	}

	return games
}


func Query_win_rate (config Settings,game uint,player_count uint) ([]Collated_player_stats, error) {
	var (
		win_rate_query *sql.Stmt
		stats Collated_player_stats
		all_stats []Collated_player_stats
		result *sql.Rows
	)
	db, err := sql.Open(config.db_driver,config.db_address)
	defer db.Close()
	Error_check(err)

	if player_count == 0 {
		win_rate_query, err = db.Prepare(`
		SELECT
		players.name_first AS name,
		CASE
			WHEN (COUNT(player_data.win) FILTER (WHERE match_data.game_id = $1)) > 0
			THEN (COUNT(player_data.win) FILTER (WHERE player_data.win = true AND match_data.game_id = $2))::float / (COUNT(player_data.win) FILTER (WHERE match_data.game_id = $3))::float
			ELSE -1
		END as win_rate,
		AVG(player_data.score) FILTER (WHERE match_data.game_id = $4) as average_score
		FROM player_data
		JOIN match_data ON match_data.id = player_data.match_id
		JOIN players ON players.id = player_data.player_id
		GROUP BY players.name_first;
		`)
		Error_check(err)
		result, err = win_rate_query.Query(game,game,game,game)

	} else {
		win_rate_query, err = db.Prepare(`
		SELECT
		players.name_first AS name,
		CASE
			WHEN (COUNT(player_data.win) FILTER (WHERE match_data.game_id = $1)) > 0
			THEN (COUNT(player_data.win) FILTER (WHERE player_data.win = true AND match_data.game_id = $2))::float / (COUNT(player_data.win) FILTER (WHERE match_data.game_id = $3))::float
			ELSE -1
		END as win_rate,
		AVG(player_data.score) FILTER (WHERE match_data.game_id = $4) as average_score
		FROM player_data
		JOIN match_data ON match_data.id = player_data.match_id
		JOIN players ON players.id = player_data.player_id
		WHERE match_data.player_count = $5
		GROUP BY players.name_first;
		`)
		Error_check(err)
		result, err = win_rate_query.Query(game,game,game,game,player_count)
	}

	defer result.Close()
	Error_check(err)

	for result.Next() {	
		result.Scan(&stats.name,&stats.win_rate,&stats.avg_score)
		all_stats = append(all_stats,stats)
	}

	return all_stats, nil
}

func Print_win_rate (all_stats []Collated_player_stats) {
	fmt.Println("Player: Win rate")
	for i := 0; i < len(all_stats); i++{
		if all_stats[i].win_rate != -1 {
			fmt.Printf("%s: %.2f%s -- %.2f points on average \n", all_stats[i].name, all_stats[i].win_rate * 100, "%",all_stats[i].avg_score)
		}
	}
}
