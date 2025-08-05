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

	s := state{
		cfg: &cfg,
	}

	c := commands{
		handlers: map[string]func(*state, command) error{
			"login": handlerLogin,
		},
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	fmt.Printf("Config: %+v\n", cfg)
}
