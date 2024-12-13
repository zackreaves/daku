package main

import (
	"fmt"
)

func new_game (name string, ties bool, tie_breakers bool, scoring bool, extensions bool) string {
	sql_stmt := fmt.Sprint("INSERT INTO games (name,ties_possible,tie_breakers,score_kept,round_extensions) VALUES (", name, ",", ties, ",", tie_breakers, ",", scoring, ",", extensions,")")
	return sql_stmt	
}

type round_data struct {
	rounds uint
	ties uint
	datetime string
	game string
}

type player_data struct {
	name string
	points float64
	win bool
	ties uint
	round_number uint
}

func Add_to_db_prompt (db_loc string) {
}
