package main

import (
	"errors"
	"fmt"

	"github.com/havokmoobii/gator/internal/config"
)

 type state struct {
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

 func handlerLogin (s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Error: Expected username.")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User", cmd.arguments[0], "has been set.")

	return nil
 }