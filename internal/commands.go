package internal

import (
	"errors"
	_ "log"

	"github.com/SyncTank/grattor/internal/database"
)

type State struct {
	Cfg  *Config
	DB   *database.Queries
	Coms *commands
}

type command struct { // ex. name = "login" , args = [username : String]
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*State, command) error
}

func (s *State) State_init(args []string) error {
	s.Coms = &commands{}
	s.Coms.registeredCommands = make(map[string]func(*State, command) error)

	cfg, err := ReadConfig()
	Check("Init - config setup", err)
	s.Cfg = &cfg

	return nil
}

func CommandSetup(args []string) command {
	return command{Name: args[1], Args: args[2:]}
}

func (c *commands) Run(s *State, cmd command) error {
	if c.registeredCommands == nil {
		return errors.New("commands map is not initialized")
	}
	fun, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New(" Run - Failed to execute command ")
	}
	return fun(s, cmd)

}

// Makes a new command ( Func_Name , Callable)
func (c *commands) Register(name string, f func(*State, command) error) {
	c.registeredCommands[name] = f
}
