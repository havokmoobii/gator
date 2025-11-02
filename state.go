package main

import (
	"errors"
	"context"

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

func (c commands) run(s *state, cmd command) error {
	_, exists := c.handlerFunctions[cmd.name]
	if !exists {
		return errors.New("Error: Command does not exist.")
	}

	return c.handlerFunctions[cmd.name](s, cmd)
}

 func (c *commands) register(name string, f func(*state, command) error) {
	c.handlerFunctions[name] = f
}

func register_commands(cmds commands) commands {
	cmds.handlerFunctions = make(map[string]func(*state, command) error)

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	return cmds
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil {
			return errors.New("Error: Must be logged in to run this command. Use 'login' or 'register'.")
		}
		return handler(s, cmd, user)
	}
}

 