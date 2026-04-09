package internal

import (
	"errors"
	"github.com/SyncTank/grattor/internal/config"
	"github.com/SyncTank/grattor/internal/database"
)

type State struct {
	Cfg  *config.Config
	db   *database.Queries
	Coms commands
}

type command struct { // ex. name = "login" , args = [username : String]
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*State, command) error
}

func commandSetup(args []string) *command {
	return &command{Name: args[1], Args: args[2:]}
}

func (c *commands) run(s *State, cmd command) error {
	if c.registeredCommands == nil {
		return errors.New("commands map is not initialized")
	}
	fun, err := c.registeredCommands[cmd.Name]
	if err {
		return errors.New(" Run - Failed to execute command ")
	}
	return fun(s, cmd)
}

// Makes a new command ( Func_Name , Callable)
func (c *commands) register(name string, f func(*State, command) error) {
	c.registeredCommands[name] = f
}
