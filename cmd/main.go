package main

import (
	"github.com/kuckjwi0928/flexi-mocker/configs"
	"github.com/kuckjwi0928/flexi-mocker/internal/server"
)

func main() {
	config := configs.NewConfig()

	s := server.NewServer(config)
	s.Run(config.Server.Port)
}
