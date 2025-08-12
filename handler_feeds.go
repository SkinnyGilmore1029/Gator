package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SkinnyGilmore1029/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeeds(s *state, cmd command) error {
	//Get all feeds from the database
	allFeeds, err := s.db.ListFeeds(context.Background())
	// Check for errors and return error if there is one
	if err != nil {
		return err
	}
	// Loop throgh allFeeds and print
	for _, feed := range allFeeds {
		fmt.Printf("Feed Name: %s\n", feed.FeedName)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("User Name: %s\n", feed.UserName)
	}
	// no errors happened dont return nil and just print
	return nil
}

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

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: newFeed.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}
	fmt.Printf("Feed created: %s\n", newFeed.Name)
	fmt.Printf("URL: %s\n", newFeed.Url)
	fmt.Printf("You are now following this feed! %s\n", newFeedFollow.FeedName)

	return nil
}
