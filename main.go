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
	//log.Println(len(args))
	if len(args) <= 1 {
		log.Fatal("Usage: <command> [arguments]")
	}
	state.State_init(os.Args)

	state.Cfg.DBString = internal.BuildDBString(&state)
	db, err := sql.Open("postgres", state.Cfg.DBString)
	if err != nil {
		log.Fatalf("Error connecting to db: %v\n", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	state.DB = dbQueries

	state.Coms.Register("login", internal.HandlerLogin)
	state.Coms.Register("register", internal.HandlerRegister)
	state.Coms.Register("reset", internal.HandlerReset)
	state.Coms.Register("users", internal.HandlerListUsers)
	state.Coms.Register("agg", internal.HandlerAggregate)

	err = state.Coms.Run(&state, internal.CommandSetup(args))
	if err != nil {
		log.Fatal(err)
	}
}
