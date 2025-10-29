package main

import (
	"fmt"
	"os"

	"github.com/havokmoobii/gator/internal/config"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: Expected command.")
		os.Exit(1)
	}

	var cliState state
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cliState.config = &cfg

	var cmds commands
	cmds.handlerFunctions = make(map[string]func(*state, command) error)
	cmds.register("login", handlerLogin)

	var cmd command
	cmd.name = os.Args[1]

	for i := 2; i < len(os.Args); i++ {
		cmd.arguments = append(cmd.arguments, os.Args[i])
	}

	err = cmds.run(&cliState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
