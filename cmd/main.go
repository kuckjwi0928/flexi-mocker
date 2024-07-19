package main

import (
	"github.com/kuckjwi0928/flexi-mocker/configs"
	"github.com/kuckjwi0928/flexi-mocker/internal"
)

func main() {
	config := configs.NewConfig()

	s := internal.NewServer(config)
	s.Run(config.Server.Port)
}
