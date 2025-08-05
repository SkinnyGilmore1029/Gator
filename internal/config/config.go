package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(config Config) error {
	// get the path
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// create or write to the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// put the info in the json file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		return err
	}
	// else return nothing
	return nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

func Read() (Config, error) {
	//get the path
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	// make sure we close the file when we are done
	defer file.Close()

	config := Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

type state struct {
	cfg *Config
}

type command struct {
	name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.handlers[cmd.Args[0]]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Args[0])
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires username")
	}
	err := s.cfg.SetUser(cmd.name)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set!")
	return nil
}
