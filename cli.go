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

func Cli () error {
	switch os.Args[1] {
	case "init":
		return init_arg(2)
	case "list":
		return list_arg(2)
	case "csv":
		csv_arg(2)
	case "tui":
		return tui_arg(2)
	default:
		fmt.Println("No argument given.") // TODO: REPLACE WITH --help TYPE OUTPUT.
	}
	return nil
}

func init_arg(arg_start_point uint) error {
		config.flags(os.Args[arg_start_point:])
		return Init(config)
}

func list_arg (arg_start_point uint) error {
	switch os.Args[arg_start_point] {
	case "players":
		config.flags(os.Args[arg_start_point+1:])
		_, col, err := Query_name(config)
		if err != nil {
			return err
		}
		fmt.Println(col) // TODO: ADD COMPONENT TO ACTUALLY PRINT PLAYER NAMES.
	case "games":
		config.flags(os.Args[arg_start_point+1:])
		Query_games(config)
	case "winrates":
		config.flags(os.Args[arg_start_point+3:])
		game, err := strconv.ParseUint(os.Args[arg_start_point+1],10,64)
		if err != nil {
			return err
		}
		player_num, err := strconv.ParseUint(os.Args[arg_start_point+2],10,64)
		if err != nil {
			return err
		}
		win_rates, err := Query_win_rate(config,uint(game),uint(player_num))
		if err != nil {
			return err
		}
		Print_win_rate(win_rates)
	}
	return nil
}

func csv_arg (arg_start_point uint) error {
	switch os.Args[arg_start_point] {
	case "table":
		config.flags(os.Args[arg_start_point+3:])
		err := Csv_insert(os.Args[arg_start_point+2],os.Args[arg_start_point+1])
		if err != nil {
			return err
		}
	case "match":
		config.flags(os.Args[arg_start_point+3:])
		matches, players, err := Match_populate(os.Args[arg_start_point+1],os.Args[arg_start_point+2])
		if err != nil {
			return err
		}
		err = Match_sort_insert(config, matches, players)
		if err != nil {
			return err
		}
	}
	return nil
}

func tui_arg (arg_start_point uint) error {
	config.flags(os.Args[arg_start_point:])
	match, err := Match_input_form(config)
	if err != nil {
		return err
	}
	_, err = fmt.Println("Game ID: ",match.game_id,"Player Count: ",match.player_count)
	return err
}
