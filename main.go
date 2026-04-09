package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/SyncTank/grattor/internal"
	"github.com/SyncTank/grattor/internal/config"
	"github.com/SyncTank/grattor/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	args := os.Args

	state := internal.State{}
	coms.registeredCommands = make(map[string]func(*State, command) error)

	Config, err := readConfig()
	Check("Init - config setup", err)
	state.Cfg = &Config

	coms.register("login", handlerLogin)
	state.Coms = coms
	if len(args) <= 2 {
		log.Panicln(" Init - Failed to capture target")
	}

	cmd := commandSetup(args)
	log.Println(coms)
	log.Println(state)
	coms.run(&state, *cmd)

	return State{}

	state.Cfg.DBString = buildDBString(&state)
	db, err := sql.Open("postgres", state.Cfg.DBString)
	Check("database connection : ", err)
	dbQueries := database.New(db)
	state.db = dbQueries

}
