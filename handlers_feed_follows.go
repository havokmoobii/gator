package main

import (
	"errors"
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"

	"github.com/havokmoobii/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected feed url.")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	feedFollowArgs := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowArgs)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has followed feed \"%s\".\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected feed url.")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	unfollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), unfollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s has unfollowed \"%s\".", user.Name, feed.Name)

	return nil
}