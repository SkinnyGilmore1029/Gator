package main

import (
	"context"
	"fmt"

	"github.com/SkinnyGilmore1029/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command) error {
	// check the lenght to make sure its valid
	if len(cmd.Args) < 2 {
		return fmt.Errorf("addfeed requires a name and a url")
	}

	// Get the current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// get and save info into variables to use
	params := database.CreateFeedParams{
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
		UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
	}

	newFeed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}
	fmt.Println(newFeed)
	return nil
}
