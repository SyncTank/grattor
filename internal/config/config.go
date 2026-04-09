package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/user"
)

const CONFIG_FILE = "/.gatorconfig.json"

type Config struct {
	DB_url            []string `json:"db_url"`
	Current_user_name string   `json:"current_user_name"`
	DBString          string
	password          string
}

func buildDBString(s *State) string {
	// protocol://username:password@host:port/database?sslmode=disable
	// postgres://Rudy:@localhost:5432/gator
	// postgres://postgres:postgres@localhost:5432/gator
	return s.Cfg.DB_url[0] + s.Cfg.Current_user_name + ":" + s.Cfg.password + "@" + "localhost:5432/gator" + s.Cfg.DB_url[1]
}

func readConfig() (Config, error) {
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
		result.SetUserConfig(cusr.Username)
	}

	return result, nil
}

func (cfg *Config) writeConfig() error {
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

func (cfg *Config) SetUserConfig(name string) error {
	cfg.Current_user_name = name
	err := cfg.writeConfig()
	log.Println("Set local state user to ", name)
	return CheckSlient(" SetUser - checking if name can be set : ", err)
}

func (cfg *Config) SetUserlocal(name string) error {
	cfg.Current_user_name = name
	log.Println("Set local state user to ", name)
	return nil
}
