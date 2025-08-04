package main

import (
	"fmt"
	"log"

	"github.com/SkinnyGilmore1029/gator/internal/config"
)

func main() {
	//Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}

	err = cfg.SetUser("Josh")
	if err != nil {
		log.Fatal("Error setting user", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	fmt.Printf("Config: %+v\n", cfg)
}
