package configs

import (
	"github.com/pelletier/go-toml"
	"os"
)

// Config is a struct for the config.toml file
type Config struct {
	Server ServerConfig  `toml:"server"`
	Routes []RouteConfig `toml:"routes"`
}

// ServerConfig is a struct for the server config in the config.toml file
type ServerConfig struct {
	Port int `toml:"port"`
}

// RouteConfig is a struct for the routes in the config.toml file
type RouteConfig struct {
	Method   string `toml:"method"`
	Path     string `toml:"path"`
	Response any    `toml:"response"`
}

func NewConfig() *Config {
	cfg := new(Config)
	configFileName := "config.toml"
	if configFilePath := os.Getenv("CONFIG_FILE_PATH"); configFilePath != "" {
		configFileName = configFilePath
	}
	bytes, err := os.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(bytes, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
