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

func (s *Settings ) flags (arguments []string) {
	var (
		arg_flags *flag.FlagSet = flag.NewFlagSet("arg_flags", flag.ExitOnError)
		db_address_override *string = arg_flags.String("a","","Override environmental variables.")
		db_env_defined bool = true 
		home,_ = os.UserHomeDir()
	)

	arg_flags.Parse(arguments)

	if *db_address_override == "" {
		s.db_address, db_env_defined = os.LookupEnv("DAKU_DB_ADDRESS")
		if db_env_defined && s.db_driver == "sqlite3" {
			s.db_address += "/DAKU.db"
		} else if s.db_driver == "sqlite3" {
			os.Setenv("DAKU_DB", home + "/.local/share/DAKU.db")
			s.db_address = os.Getenv("DAKU_DB")
		}
	} else {
		s.db_address = *db_address_override
		fmt.Println("Set DAKU_DB in your shell environment.")
	}
}

func (s *Settings) driver_chooser () {
	s.db_driver = os.Getenv("DAKU_SQL_SERVICE")
	if s.db_driver == "" {
		s.db_driver = "sqlite3"
	}
}

var config Settings

func main () {
	config.driver_chooser()
	switch os.Args[1] {
	case "init":
		config.flags(os.Args[2:])
		Init(config.db_driver,config.db_address)
		fmt.Println(config.db_address)
	case "query":
		config.flags(os.Args[3:])
		res := Query(config.db_address, os.Args[2])
		fmt.Println(res)
	case "csv":
		config.flags(os.Args[4:])
		switch os.Args[2] {
		case "players":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Players{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp) //FIXME AND ALL LIKE INSTANCES: INTERFACE UNNECESSARY AND RESULTING FUNCTIONS ARE UNNECESSARY OR NEED A REVISED IMPLEMENTATION.
				Insert_from_table(config.db_driver,config.db_address,tp)
			}
		case "games":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Games{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp)
				Insert_from_table(config.db_driver,config.db_address,tp)
			}
		case "match_data":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Match_data{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp)
				Insert_from_table(config.db_driver,config.db_address,tp)
			}
		case "player_data":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Player_data{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp)
				Insert_from_table(config.db_driver,config.db_address,tp)
			}
		}
	default:
		fmt.Println("No argument given.")
	}
}
