package main

import (
	"os"
	"flag"
)

var db_location = flag.String("l","$HOME/.local/share/DAKU.db","Override location for database.")
func main () {
	flag.Parse()
	db_env, dbEnvDefined := os.LookupEnv("DAKU_DB")
	if dbEnvDefined == false {
		os.Setenv("DAKU_DB",*db_location)
		db_env = os.Getenv("DAKU_DB")
	}
	Init(db_env)
}
