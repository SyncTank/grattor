package internal

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/SyncTank/grattor/internal/database"
)

func HandlerLogin(s *State, cmd command) error {
	if len(cmd.Args) != 1 { // expects single arg, username
		log.Fatalln(" Run - Failed to execute command ")
		return errors.New(" Handler expects a single argument, the username")
	}

	com_name := cmd.Args[0]
	usr, err := s.DB.GetUser(context.Background(), com_name)
	if err != nil {
		log.Panicln("User not found")
		os.Exit(1)
	}

	erro := s.Cfg.SetUserConfig(usr.Name)
	if erro != nil {
		return errors.New(" failed to set user")
	}

	return nil
}

func printUser(user database.User) {
	log.Printf(" * ID: %v\n", user.ID)
	log.Printf(" * Name: %v\n", user.Name)
}
