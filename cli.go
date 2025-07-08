package main

import (
	"flag"
	"fmt"
	"os"
)

type Settings struct {
	db_address string
	db_driver string
}

func (s *Settings ) flags (arguments []string) (error) {
	var (
		arg_flags *flag.FlagSet = flag.NewFlagSet("arg_flags", flag.ExitOnError)
		db_address_override *string = arg_flags.String("a","","Override environmental variables.")
	)


	err := arg_flags.Parse(arguments)

	if *db_address_override == "" {
		s.db_address, _ = os.LookupEnv("DAKU_DB_ADDRESS")
	} else {
		s.db_address = *db_address_override
		fmt.Println("Set DAKU_DB in your shell environment.")
	}

	return err
}

func list_flags (arguments []string) (uint, uint, int, error) {
	var (
		list_flags *flag.FlagSet = flag.NewFlagSet("list_flags", flag.ExitOnError)
		player_count uint
		game_id uint
		round int
		db_address_override *string = list_flags.String("a","","Override environmental variable.")
	)

	list_flags.UintVar(&game_id, "g", 0, "Enter game id.")
	list_flags.UintVar(&player_count, "c", 0, "Enter number of players")
	list_flags.IntVar(&round,"r",1, "Enter which round is of interest, 0 returns the aggregate and anything less returns the last round's values.")

	err := list_flags.Parse(arguments)

	if *db_address_override == "" {
		config.db_address, _ = os.LookupEnv("DAKU_DB_ADDRESS")
	} else {
		config.db_address = *db_address_override
		fmt.Println("Set DAKU_DB in your shell environment.")
	}

	return game_id, player_count, round, err
}

var config = Settings {
	db_driver: "postgres",
}

func Cli () error {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "init":
			return init_arg(2)
		case "list":
			return list_arg(2)
		case "csv":
			return csv_arg(2)
		case "tui":
			return tui_arg(2)
			}
		} else {
			fmt.Println("No argument given.")
			fmt.Println("init - initialize database tables.")
			fmt.Println("list {players,games,winrates game.id match_data.player_count} - lists players, games, and winrates")
			fmt.Println("csv {table {players, games, match_data, player_data} csv_path, match match_data_path player_data_path} - imports data from csv, best to use match option.")
			fmt.Println("tui - starts incomplete tui form.")
		}
	
	return nil
}

func init_arg (arg_start_point uint) error {
		config.flags(os.Args[arg_start_point:])
		return Init(config)
}

func list_arg (arg_start_point uint) error {
	switch os.Args[arg_start_point] {
	case "players":
		config.flags(os.Args[arg_start_point+1:])
		player_list, _, err := Query_name(config)
		if err != nil {
			return err
		}
		Print_player_list (player_list)
	case "games":
		config.flags(os.Args[arg_start_point+1:])
		game_list, _, err := Query_games(config)
		if err != nil {
			return err
		}
		Print_game_list(game_list)
	case "winrates":
		game_id, player_num, round_num, _ := list_flags(os.Args[arg_start_point+1:])
		win_rates, err := Query_win_rate(config,game_id,player_num,round_num)
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
