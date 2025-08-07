package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires username")
	}

	// Get users name
	name := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Println("User does not exist")
		os.Exit(1)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return err
	}
	fmt.Println("User has been set!")
	return nil
}
