package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Settings struct {
	db_address string
	db_driver string
}

func (s *Settings ) flags (arguments []string) {
	var (
		arg_flags *flag.FlagSet = flag.NewFlagSet("arg_flags", flag.ExitOnError)
		db_address_override *string = arg_flags.String("a","","Override environmental variables.")
	)

	arg_flags.Parse(arguments)

	if *db_address_override == "" {
		s.db_address, _ = os.LookupEnv("DAKU_DB_ADDRESS")
	} else {
		s.db_address = *db_address_override
		fmt.Println("Set DAKU_DB in your shell environment.")
	}
}

func (s *Settings) driver_chooser () {
	s.db_driver = "postgres"
}

var config Settings

func main () {
	config.driver_chooser()
	switch os.Args[1] {
	case "init":
		config.flags(os.Args[2:])
		Error_check(Init(config))
	case "csv":
		switch os.Args[2] {
		case "table":
			config.flags(os.Args[5:])
			Error_check(Csv_insert(os.Args[4],os.Args[3]))
		case "match":
			config.flags(os.Args[5:])
			matches, players := Match_populate(os.Args[3],os.Args[4])
			Match_sort_insert(config, matches, players)
		}
	case "list":
		switch os.Args[2] {
		case "players":
			config.flags(os.Args[3:])
			_, col, err := Query_name(config)
			Error_check(err)
			fmt.Println(col) // TODO: ADD COMPONENT TO ACTUALLY PRINT PLAYER NAMES.
		case "games":
			config.flags(os.Args[3:])
			Query_games(config)
		case "winrates":
			config.flags(os.Args[5:])
			game, err := strconv.ParseUint(os.Args[3],10,64)
			Error_check(err)
			player_num, err := strconv.ParseUint(os.Args[4],10,64)
			Error_check(err)
			win_rates, err := Query_win_rate(config,uint(game),uint(player_num))
			Error_check(err)
			Print_win_rate(win_rates)
		}
	case "tui":
		config.flags(os.Args[2:])
		match := Match_input_form(config)
		fmt.Println("Game ID: ",match.game_id,"Player Count: ",match.player_count)
	default:
		fmt.Println("No argument given.")
	}
}
