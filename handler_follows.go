package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/SkinnyGilmore1029/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	// make sure they actually typed something
	if len(cmd.Args) < 1 {
		return fmt.Errorf("please enter a url you want to follow")
	}

	// get the url from the command line
	url := cmd.Args[0]
	feedID, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		fmt.Println("url is not valid")
		os.Exit(1)
	}

	//get the current user that is logged in
	current_user_name, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	result, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    uuid.NullUUID{UUID: current_user_name.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feedID, Valid: true},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Name of feed %s\n", result.FeedName)
	fmt.Printf("Name of current user %s\n", result.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	//get current user
	current_user_name, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// get the feeds the user follows
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), uuid.NullUUID{UUID: current_user_name.ID, Valid: true})
	if err != nil {
		return err
	}

	//need a for loop to go through current user name follows?
	for _, follow := range follows {
		fmt.Printf("Feed name: %s\n", follow.FeedName)
	}

	return nil

}
