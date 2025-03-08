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
		Init(config)
		fmt.Println(config.db_address)
	case "csv":
		config.flags(os.Args[4:])
		Csv_insert(os.Args[3],os.Args[2])
	case "list":
		switch os.Args[2] {
		case "players":
			config.flags(os.Args[3:])
			_, col := Query_name(config)
			fmt.Println(col)
		case "games":
			config.flags(os.Args[3:])
			Query_games(config)
		}
	default:
		fmt.Println("No argument given.")
	}
}
