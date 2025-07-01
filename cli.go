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

var config = Settings {
	db_driver: "postgres",
}

func Cli () {
	switch os.Args[1] {
	case "init":
		init_arg(2)
	case "list":
		list_arg(2)
	case "csv":
		csv_arg(2)
	case "tui":
		tui_arg(2)
	default:
		fmt.Println("No argument given.")
	}
}

func init_arg(arg_start_point uint) {
		config.flags(os.Args[arg_start_point:])
		Error_check(Init(config))
}

func list_arg (arg_start_point uint) {
	switch os.Args[arg_start_point] {
	case "players":
		config.flags(os.Args[arg_start_point+1:])
		_, col, err := Query_name(config)
		Error_check(err)
		fmt.Println(col) // TODO: ADD COMPONENT TO ACTUALLY PRINT PLAYER NAMES.
	case "games":
		config.flags(os.Args[arg_start_point+1:])
		Query_games(config)
	case "winrates":
		config.flags(os.Args[arg_start_point+3:])
		game, err := strconv.ParseUint(os.Args[arg_start_point+1],10,64)
		Error_check(err)
		player_num, err := strconv.ParseUint(os.Args[arg_start_point+2],10,64)
		Error_check(err)
		win_rates, err := Query_win_rate(config,uint(game),uint(player_num))
		Error_check(err)
		Print_win_rate(win_rates)
	}
}

func csv_arg (arg_start_point uint) {
	switch os.Args[arg_start_point] {
	case "table":
		config.flags(os.Args[arg_start_point+3:])
		Error_check(Csv_insert(os.Args[arg_start_point+2],os.Args[arg_start_point+1]))
	case "match":
		config.flags(os.Args[arg_start_point+3:])
		matches, players := Match_populate(os.Args[arg_start_point+1],os.Args[arg_start_point+2])
		Match_sort_insert(config, matches, players)
	}
}

func tui_arg (arg_start_point uint) {
	config.flags(os.Args[arg_start_point:])
	match := Match_input_form(config)
	fmt.Println("Game ID: ",match.game_id,"Player Count: ",match.player_count)
}
