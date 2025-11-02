package main

import (
	"errors"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"

	"github.com/havokmoobii/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected feed name and url.")
	}

	if len(cmd.arguments) < 2 {
		return errors.New("Error: Expected feed url.")
	}

	feedArgs := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.arguments[0],
		Url: cmd.arguments[1],
		UserID: user.ID,
	}

	_, err := s.db.CreateFeed(context.Background(), feedArgs)
	if err != nil {
		return err
	}

	fmt.Printf("Feed \"%s\" created.\n", feedArgs.Name)

	// Follow expects a url argument, which is the second argument for addfeed.
	cmd.arguments[0] = cmd.arguments[1]
	err = handlerFollow(s, cmd, user)
	if err != nil {
		return err
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		username, err := s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name, "-", feed.Url, "-", username)
	}

	return nil
}