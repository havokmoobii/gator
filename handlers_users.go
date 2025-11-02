package main

import (
	"errors"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"

	"github.com/havokmoobii/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected username.")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	err = s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User", cmd.arguments[0], "has been set.")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected username.")
	}

	userArgs := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.arguments[0],
	}

	_, err := s.db.CreateUser(context.Background(), userArgs)
	if err != nil {
		return err
	}

	err = s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User", userArgs.Name, "created.")

	return nil
}

func handlerUsers(s *state, cmd command) error {
	usernames, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, username := range usernames {
		if username == s.config.Current_user_name {
			fmt.Println(username, "(current)")
			continue
		}
		fmt.Println(username)
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Database reset!")

	return nil
}