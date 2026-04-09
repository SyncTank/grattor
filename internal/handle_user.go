package internal

import (
	"errors"
	"log"
)

func HandlerLogin(s *State, cmd command) error {
	if len(cmd.Args) != 1 { // expects single arg, username
		log.Fatalln(" Run - Failed to execute command ")
		return errors.New(" Handler expects a single argument, the username")
	}
	com_name := cmd.Args[0]

	err := s.Cfg.SetUserConfig(com_name)
	if err != nil {
		return errors.New(" failed to set user")
	}

	return nil
}
