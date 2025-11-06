package internal

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const CONFIG_FILE = "/.gatorconfig.json"

type Config struct {
	DB_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func ReadConfig() Config {
	var result Config

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err = json.Unmarshal(data, &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func write(cfg Config) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

}

func SetUser(name string, cfg Config) {
	cfg.Current_user_name = name
	write(cfg)
}
