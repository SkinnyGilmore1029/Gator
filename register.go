package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SkinnyGilmore1029/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	//check to make sure they gave a name
	if len(cmd.Args) == 0 {
		return errors.New("please type a name to register")
	}

	//get info from command line
	name := cmd.Args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	// lol what do i do with you
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		// Check if the error indicates the user already exists
		// This might be a constraint violation or duplicate key error
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique constraint") {
			fmt.Println("User already exists")
			os.Exit(1) // Exit with code 1 as required
		}
		// If it's some other database error, return it
		return err
	}

	s.cfg.CurrentUserName = name

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s created successfully!\n", name)
	fmt.Printf("User data: %+v\n", user)

	return nil
}
