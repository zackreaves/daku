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
		db_override *string = arg_flags.String("db","","Override environmental variables.")
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
		Init(db_loc)
		fmt.Println(db_loc)
	case "query":
		db_loc := arg_flags(os.Args[3:])
		Exec(db_loc, os.Args[2])
	case "sqlite":
		db_loc := arg_flags(os.Args[2:])
		exec.Command("sqlite3",db_loc)
	default:
		fmt.Println("No argument given.")
	}
}
