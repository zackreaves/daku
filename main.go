package main

import (
	"os"
	"flag"
	"fmt"
)

func arg_flags (arguments []string) string {
	var arg_flags *flag.FlagSet = flag.NewFlagSet("arg_flags", flag.ExitOnError)
	var db_override *string = arg_flags.String("db","","Override environmental variables.")
	var db_env string = ""
	var db_env_defined = true
	arg_flags.Parse(arguments)

	if *db_override == "" {
		db_env, db_env_defined = os.LookupEnv("DAKU_DB")
		if db_env_defined == false {
			os.Setenv("DAKU_DB",os.Getenv("HOME") + "/.local/share/DAKU.db")
			db_env = os.Getenv("DAKU_DB")
		} 
	} else {
		db_env = *db_override
		fmt.Println("Set your DAKU_DB in your shell environment.")
	}

	return db_env
}

func main () {
	db_loc := arg_flags(os.Args[2:])
	switch os.Args[1] {
	case "init":
		Init(db_loc)
		fmt.Println(db_loc)
	default:
		fmt.Println("No argument given.")
	}
	fmt.Println(os.Args)
}
