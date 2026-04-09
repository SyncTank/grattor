package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/SyncTank/grattor/internal"
	"github.com/SyncTank/grattor/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	state := internal.State{}
	args := os.Args
	if len(args) <= 2 {
		log.Fatal("Usage: <command> [arguments]")
	}
	state.State_init(os.Args)

	state.Coms.Register("login", internal.HandlerLogin)
	cmd := internal.CommandSetup(args)
	state.Coms.Run(&state, *cmd)

	state.Cfg.DBString = internal.BuildDBString(&state)
	db, err := sql.Open("postgres", state.Cfg.DBString)
	internal.Check("database connection : ", err)

	dbQueries := database.New(db)
	state.DB = dbQueries

}
