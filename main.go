package main

import (
	"flag"
	"fmt"
	"os"
)

func arg_flags (arguments []string) string {
	var (
		arg_flags *flag.FlagSet = flag.NewFlagSet("arg_flags", flag.ExitOnError)
		db_override *string = arg_flags.String("d","","Override environmental variables.")
		db_env string = ""
		db_env_defined bool = true 
		home,_ = os.UserHomeDir()
	)
		arg_flags.Parse(arguments)

	if *db_override == "" {
		db_env, db_env_defined = os.LookupEnv("DAKU_DB_ADDRESS")
		if db_env_defined {
			db_env += "/DAKU.db"
		} else {
			os.Setenv("DAKU_DB", home + "/.local/share/DAKU.db")
			db_env = os.Getenv("DAKU_DB")
		} 
	} else {
		db_env = *db_override
		fmt.Println("Set DAKU_DB in your shell environment.")
	}

	return db_env
}
func sql_chooser () string {
	sql_choice := os.Getenv("DAKU_SQL_SERVICE")

	if sql_choice == "" {
		sql_choice = "sqlite3"
	}

	return sql_choice
}

func main () {
	db_driver := sql_chooser()
	switch os.Args[1] {
	case "init":
		db_loc := arg_flags(os.Args[2:])
		Init(db_driver,db_loc)
		fmt.Println(db_loc)
	case "query":
		db_loc := arg_flags(os.Args[3:])
		res := Query(db_loc, os.Args[2])
		fmt.Println(res)
	case "csv":
		db_loc := arg_flags(os.Args[4:])
		switch os.Args[2] {
		case "players":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Players{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp) //FIXME AND ALL LIKE INSTANCES: INTERFACE UNNECESSARY AND RESULTING FUNCTIONS ARE UNNECESSARY OR NEED A REVISED IMPLEMENTATION.
				Insert_from_table(db_driver,db_loc,tp)
			}
		case "games":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Games{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp)
				Insert_from_table(db_driver,db_loc,tp)
			}
		case "match_data":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Match_data{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp)
				Insert_from_table(db_driver,db_loc,tp)
			}
		case "player_data":
			csv_arr, rows := Import_from_csv(os.Args[3])
			format := csv_arr[0]
			csv_args := csv_arr[1:]
			t := Player_data{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				Populate_from_arguments(csv_args[i], format, tp)
				Insert_from_table(db_driver,db_loc,tp)
			}
		}
	default:
		fmt.Println("No argument given.")
	}
}
