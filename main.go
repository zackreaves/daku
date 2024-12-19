package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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
		db_env, db_env_defined = os.LookupEnv("DAKU_DB")
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

func main () {
	switch os.Args[1] {
	case "init":
		db_loc := arg_flags(os.Args[2:])
		Init("sqlite3",db_loc)
		fmt.Println(db_loc)
	case "query":
		db_loc := arg_flags(os.Args[3:])
		res := Query(db_loc, os.Args[2])
		fmt.Println(res)
	case "sqlite":
		db_loc := arg_flags(os.Args[2:])
		res := exec.Command("sqlite3",db_loc)
		fmt.Println(res)
	case "ngame": //FIXME: Prompt doesn't work.
		db_loc := arg_flags(os.Args[2:])
		Add_to_db_prompt(db_loc)
	case "csv":
		db_loc := arg_flags(os.Args[4:])
		switch os.Args[2] {
		case "players":
			fmt.Println(os.Args)
			csv_arr, rows := Import_from_csv(os.Args[3])
			fmt.Println(csv_arr, db_loc)
			format := csv_arr[0][:]
			csv_args := csv_arr[1:][:]
			t := Players{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				args := csv_args[i][:]
				Populate_from_arguments(args, format, tp)
				Insert_from_table("sqlite3",db_loc,tp)
			}
		case "games":
			fmt.Println(os.Args)
			csv_arr, rows := Import_from_csv(os.Args[3])
			fmt.Println(csv_arr, db_loc)
			format := csv_arr[0][:]
			csv_args := csv_arr[1:][:]
			t := Games{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				args := csv_args[i][:]
				Populate_from_arguments(args, format, tp)
				Insert_from_table("sqlite3",db_loc,tp)
			}
		case "match_data":
			fmt.Println(os.Args)
			csv_arr, rows := Import_from_csv(os.Args[3])
			fmt.Println(csv_arr, db_loc)
			format := csv_arr[0][:]
			csv_args := csv_arr[1:][:]
			t := Match_data{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				args := csv_args[i][:]
				Populate_from_arguments(args, format, tp)
				Insert_from_table("sqlite3",db_loc,tp)
			}
		case "player_data":
			fmt.Println(os.Args)
			csv_arr, rows := Import_from_csv(os.Args[3])
			fmt.Println(csv_arr, db_loc)
			format := csv_arr[0][:]
			csv_args := csv_arr[1:][:]
			t := Player_data{}
			tp := &t
			for i := 0; i < rows-1 ; i++ {
				args := csv_args[i][:]
				Populate_from_arguments(args, format, tp)
				Insert_from_table("sqlite3",db_loc,tp)
			}
		}
	default:
		fmt.Println("No argument given.")
	}
}
