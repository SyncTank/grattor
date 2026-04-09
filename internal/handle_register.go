package internal

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/SyncTank/grattor/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd command) error {
	t := time.Now()
	args_handle := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: t,
		UpdatedAt: t,
		Name:      cmd.Name,
	}
	usr, err := s.DB.CreateUser(context.Background(), args_handle)
	if err != nil {
		os.Exit(1)
	}

	s.Cfg.SetUserConfig(usr.Name)
	log.Println(" User was created ")
	return nil
}
