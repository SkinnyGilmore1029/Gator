package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list users: %w", err)
	}

	currentUser := s.cfg.CurrentUserName

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
