package main

import (
	"context"
	"fmt"
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

/*

Assignment
Add a new feeds handler. It takes no arguments and prints all the feeds in the database to the console. Be sure to include:

The name of the feed
The URL of the feed
The name of the user that created the feed (you might need a new SQL query)


*/
