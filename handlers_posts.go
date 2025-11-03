package main

import (
	"errors"
	"context"
	"time"
	"github.com/google/uuid"
	"strings"
	"strconv"
	"fmt"

	"github.com/havokmoobii/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected time between requests.")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(feedToFetch.Url)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {
		pubTime, err := fuzzyTimeParse(item.PubDate)
		if err != nil {
			return err
		}

		postParams := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: item.Description,
			PublishedAt: pubTime,
			FeedID: feedToFetch.ID,
		}

		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if !strings.Contains(err.Error(), "violates unique constraint") {
				return err
			}
		}
	}

	return nil
}

func fuzzyTimeParse(pubDate string) (time.Time, error) {
	layouts := []string{
			time.RFC1123Z,                 // "Mon, 02 Jan 2006 15:04:05 -0700"
			time.RFC1123,                  // "Mon, 02 Jan 2006 15:04:05 MST"
			"Mon, _2 Jan 2006 15:04:05 -0700", // non–zero-padded day
			"Mon, _2 Jan 2006 15:04:05 MST",   // non–zero-padded day, named zone
			time.RFC850,                   // "Monday, 02-Jan-06 15:04:05 MST"
			time.ANSIC,                    // "Mon Jan _2 15:04:05 2006"
		}

		var pubTime time.Time
		var err error
		for _, layout := range layouts {
			pubTime, err = time.Parse(layout, pubDate)
			if err == nil {
				break
			}
		}
		if err != nil {
			return time.Now(), err
		}
		return pubTime, nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	browseLimit := 2
	if len(cmd.arguments) > 0 {
		arg, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
		browseLimit = arg
	}

	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(browseLimit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), getPostsParams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
		fmt.Println()
	}

	return nil
}