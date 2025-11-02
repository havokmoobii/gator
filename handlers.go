package main

import (
	"context"
)

// Will probably sort into a handlers_rss or something later when functionality is added.
func handlerAgg(s *state, cmd command) error {
	fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	return nil
}

