package internal

import (
	_ "context"
	_ "log"
	_ "time"
)

func HandlerAggregate(s *State, cmd command) error {
	fetchFeed(, "https://www.wagslane.dev/index.xml")
	return nil
}
