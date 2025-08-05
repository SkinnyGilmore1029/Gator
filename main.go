package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/SkinnyGilmore1029/gator/internal/config"
)

type state struct {
	cfg *config.Config
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
