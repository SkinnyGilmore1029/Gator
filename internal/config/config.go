package config

import (
	"encoding/json"
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
