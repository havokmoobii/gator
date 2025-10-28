package main

import (
	"fmt"

	"github.com/havokmoobii/gator/internal/config"
)

func main() {

	configFile, err := config.Read()

	fmt.Println(configFile, err)

	configFile.SetUser("Lane")

	fmt.Println(configFile, err)

	configFile, err = config.Read()

	fmt.Println(configFile, err)
}
