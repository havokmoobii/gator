package main

import (
	"fmt"

	"github.com/havokmoobii/gator/internal/config"
)

func main() {

	config, err := config.Read()

	fmt.Println(config, err)
}
