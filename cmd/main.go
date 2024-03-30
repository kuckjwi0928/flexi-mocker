package main

import (
	"github.com/kuckjwi0928/flexi-mocker/configs"
	"github.com/kuckjwi0928/flexi-mocker/pkg"
)

func main() {
	config := configs.NewConfig()

	s := pkg.NewServer(config)
	s.Run(config.Server.Port)
}
