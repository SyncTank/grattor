package internal

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"os/user"
)

const CONFIG_FILE = "/.gatorconfig.json"

type State struct {
	Cfg  *config
	Coms commands
}

type command struct { // ex. name = "login" , args = [username : String]
	name string
	args []string
}

type commands struct {
	map_fun map[string]func(*State, command) error
}

type config struct {
	DB_url            []string `json:"db_url"`
	Current_user_name string   `json:"current_user_name"`
	DBString          string
	password          string
}

func Check(errorContext string, err error) {
	if err != nil {
		log.Fatalln(" "+errorContext+" ", err)
		panic(err)
	}
}

func CheckSlient(errorContext string, err error) error {
	if err != nil {
		log.Fatalln(" "+errorContext+" ", err)
		return err
	}
	return nil
}

func Init(args []string) State {
	var state State
	coms := commands{make(map[string]func(*State, command) error)}

	config, err := readConfig()
	Check("Init - config setup", err)
	state.Cfg = config

	coms.register("login", handlerLogin)

	if len(args) <= 2 {
		log.Panicln(" Init - Failed to capture target")
	}

	coms.run(&state, commandSetup(args))
	state.Coms = coms

	state.Cfg.DBString = buildDBString(&state)

	return state
}

func buildDBString(s *State) string {
	// protocol://username:password@host:port/database?sslmode=disable
	// postgres://Rudy:@localhost:5432/gator
	// postgres://postgres:postgres@localhost:5432/gator
	return s.Cfg.DB_url[0] + s.Cfg.Current_user_name + ":" + s.Cfg.password + "@" + "localhost:5432/gator" + s.Cfg.DB_url[1]
}

func commandSetup(args []string) command {
	var cmd command
	cmd.name = args[1]
	cmd.args = args[2:]
	return cmd
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 { // expects single arg, username
		log.Fatalln(" Run - Failed to execute command ")
		return errors.New(" Handler expects a single argument, the username")
	}

	SetUserConfig(cmd.args[0], s.Cfg)

	return nil
}

func (c *commands) run(s *State, cmd command) error {
	err := c.map_fun[cmd.name](s, cmd)
	if err != nil {
		log.Fatalln(" Run - Failed to execute command ", err)
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*State, command) error) {
	if c.map_fun[name] == nil {
		c.map_fun[name] = f
	}
	c.map_fun[name] = f
}

func readConfig() (*config, error) {
	var result config

	dir, err := os.Getwd()
	Check("Read - Failed to fetch working directory :", err)

	file, err := os.Open(dir + CONFIG_FILE)
	CheckSlient(" Read - Failed to find file : ", err)
	defer file.Close()

	data, err := io.ReadAll(file)
	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(" Data Failed to Unmarshal : ", err)
	}

	cusr, err := user.Current()
	Check(" Failed to fetch user : ", err)

	if result.Current_user_name == "" {
		SetUserConfig(cusr.Username, &result)
	}

	return &result, nil
}

func writeConfig(cfg *config) error {
	//dir, err := os.Getwd()
	dir, err := os.UserHomeDir()
	Check(" Write - Failed to fetch working directory : ", err)

	filePath := dir + CONFIG_FILE

	data, err := json.Marshal(cfg)
	Check(" Marshal Failure : ", err)

	err = os.WriteFile(filePath, data, 0644) // replaces file
	Check(" Write out failed : ", err)

	return nil
}

func SetUserConfig(name string, cfg *config) error {
	cfg.Current_user_name = name
	err := writeConfig(cfg)
	log.Println("Set local state user to ", name)
	return CheckSlient(" SetUser - checking if name can be set : ", err)
}

func SetUser(name string, cfg *config) {
	cfg.Current_user_name = name
	log.Println("Set local state user to ", name)
}
