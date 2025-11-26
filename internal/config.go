package internal

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"io"
	"log"
	"os"
	"os/user"
)

const CONFIG_FILE = "/.gatorconfig.json"

type State struct {
	cfg *Config
}

type Command struct {
	name string
	args []string
}

type Config struct {
	DB_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Check(errorContext string, err error) {
	if err != nil {
		log.Fatalln(" "+errorContext+" ", err)
		panic(err)
	}
}

func CheckSlient(errorContext string, err error) {
	if err != nil {
		log.Fatalln(" "+errorContext+" ", err)
	}
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(" Handler expects a single argument, the username")
	}

	SetUser("Jeff", *s.cfg)

	return nil
}

func ReadConfig() (Config, error) {
	var result Config

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
		SetUser(cusr.Username, result)
	}

	return result, nil
}

func writeConfig(cfg Config) error {
	dir, err := os.Getwd()
	Check(" Write - Failed to fetch working directory : ", err)

	filePath := dir + CONFIG_FILE

	data, err := json.Marshal(cfg)
	CheckSlient(" Marshal Failure : ", err)

	err = os.WriteFile(filePath, data, 0644) // replaces file
	Check(" Write out failed : ", err)

	return nil
}

func SetUser(name string, cfg Config) {
	cfg.Current_user_name = name
	writeConfig(cfg)
}
