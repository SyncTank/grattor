package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
)

const CONFIG_FILE = "/.gatorconfig.json"

type Config struct {
	DB_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func ReadConfig() (Config, error) {
	var result Config

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(dir + CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err = json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	cusr, err := user.Current()
	if err != nil {
		log.Fatalln(" Failed to fetch user : ", err)
	}

	if result.Current_user_name == "" {
		SetUser(cusr.Username, result)
	}

	return result, nil
}

func write(cfg Config) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Open(dir + CONFIG_FILE)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, err := json.Marshal(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	n, err := file.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Bytes writen %d", n)

	return nil
}

func SetUser(name string, cfg Config) {
	cfg.Current_user_name = name
	write(cfg)
}
