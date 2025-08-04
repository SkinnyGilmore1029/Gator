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

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

/*

This package should have the following functionality exported so the main package can use it:

Export a Config struct that represents the JSON file structure, including struct tags.
Export a Read function that reads the JSON file found at ~/.gatorconfig.json and returns a Config struct. It should read the file from the HOME directory, then decode the JSON string into a new Config struct. I used os.UserHomeDir to get the location of HOME.
Export a SetUser method on the Config struct that writes the config struct to the JSON file after setting the current_user_name field.
I also wrote a few non-exported helper functions and added a constant to hold the filename.

getConfigFilePath() (string, error)
write(cfg Config) error
const configFileName = ".gatorconfig.json"
But you can implement the internals of the package however you like.

Update the main function to:
Read the config file.
Set the current user to "lane" (actually, you should use your name instead) and update the config file on disk.
Read the config file again and print the contents of the config struct to the terminal.
*/
