package internal

import (
	"context"
	"log"
	"time"
)

func HandlerListUsers(s *State, cmd command) error {
	t := time.Now()
	usrs, err := s.DB.GetUsers(context.Background())
	if err != nil {
		log.Println("Could not list users", t)
		return err
	}

	for index := range usrs {
		if usrs[index].Name == s.Cfg.Current_user_name {
			log.Println(usrs[index].Name + " (current)")
		} else {
			log.Println(usrs[index].Name)
		}
	}

	return nil
}
