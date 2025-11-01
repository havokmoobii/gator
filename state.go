package main

import (
	"errors"

	"github.com/havokmoobii/gator/internal/database"
	"github.com/havokmoobii/gator/internal/config"
)

type state struct {
	db  *database.Queries
	config *config.Config
}

type command struct {
	name string
	arguments []string
}

type commands struct {
	handlerFunctions map[string]func(*state, command) error
}

func (c commands) run (s *state, cmd command) error {
	_, exists := c.handlerFunctions[cmd.name]
	if !exists {
		return errors.New("Error: Command does not exist.")
	}

	return c.handlerFunctions[cmd.name](s, cmd)
}

 func (c *commands) register(name string, f func(*state, command) error) {
	c.handlerFunctions[name] = f
}

 