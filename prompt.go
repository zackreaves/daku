package main

import (
	"fmt"
)

func new_game (name string, ties uint8, tie_breakers uint8, scoring uint8, extensions uint8) string {
	sql_stmt := fmt.Sprint("INSERT INTO games (name,ties_possible,tie_breakers,score_kept,round_extensions) VALUES (", name, ",", ties, ",", tie_breakers, ",", scoring, ",", extensions,")")
	return sql_stmt	
}

func yn_to_bool (ans string) (uint8, error) { // Create a 1/0 bool for SQL.
	var (
		err error
		ans_conv uint8
	)
	if ans == "y" || ans == "Y" || ans == "yes" || ans == "Yes" {
		ans_conv = 1
	} else if ans == "n" || ans == "N"{
		ans_conv = 0
	} else {
		err = fmt.Errorf("Invalid answer.")
	}
	return ans_conv, err
}

func Add_to_db_prompt (db_loc string) {
	var (
		purpose string
		sql_stmt string
	)
	
	fmt.Scanln(purpose)
	switch purpose {
	case "game":
		var (
			name string
			ties_possible string
			tie_breakers string
			scoring string
			extensions string
		)
		fmt.Println("Name the Game")
		_,err := fmt.Scanln(&name)
		Error_check(err)
		fmt.Println("Does it allow for ties in any way? [y/n]")
		_,err = fmt.Scanln(&ties_possible)
		Error_check(err)
		ties_bool, err := yn_to_bool(ties_possible)
		Error_check(err)
		fmt.Println("Does the game have built in tie breaking? [y/n]")
		_,err = fmt.Scanln(&tie_breakers)
		Error_check(err)
		tie_break_bool, err := yn_to_bool(tie_breakers)
		Error_check(err)
		fmt.Println("Is there scoring in the game?")
		_,err = fmt.Scanln(&scoring)
		Error_check(err)
		score_bool, err := yn_to_bool(scoring)
		Error_check(err)
		fmt.Println("Does the game allow for extensions on existing matches? [y/n]")
		_,err = fmt.Scanln(&extensions)
		Error_check(err)
		extensions_bool, err := yn_to_bool(extensions)

		sql_stmt = new_game(name,ties_bool,tie_break_bool,score_bool,extensions_bool)
		Exec(db_loc, sql_stmt)
	default:
		err := fmt.Errorf("No arguments given.")
		Error_check(err)
	}
}
