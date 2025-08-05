package main

import (
	"errors"
	"fmt"
	"log"
	"os"

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
	if handler, ok := c.handlers[cmd.name]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires username")
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set!")
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
		handlers: make(map[string]func(*state, command) error),
	}

	c.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	commandName := os.Args[1]
	arguments := os.Args[2:]

	cmd := command{
		name: commandName,
		Args: arguments,
	}

	if err := c.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

/*


cfg, err = config.Read()
	if err != nil {
		log.Fatal("Error reading config file", err)
	}
	fmt.Printf("Config: %+v\n", cfg)

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args)

    if len(os.Args) < 2 {
        fmt.Println("not enough arguments")
        return
    }

    commandName := os.Args[1] // "greet"
    arguments := os.Args[2:]  // ["boots"]

    fmt.Printf("Command: %s\n", commandName)
    fmt.Printf("Arguments: %v\n", arguments)
}
*/
