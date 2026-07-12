package internal

import (
	"context"
	"log"
	"time"
)

func HandlerReset(s *State, cmd command) error {
	t := time.Now()
	err := s.DB.Reset(context.Background())
	if err != nil {
		log.Println("Reset could not be done")
		return err
	}

	log.Println("Reset completed at : ", t)
	return nil
}
