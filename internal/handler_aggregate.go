package internal

import (
	"context"
	"log"
)

func HandlerAggregate(s *State, cmd command) error {
	log.Printf("state %v\n", s)
	log.Printf("cmd %v\n", cmd)
	fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	return nil
}
