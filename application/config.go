package application

import (
	"os"
	"strconv"
)

// configuration structure
type Config struct {
	ServerPort uint16
	Environment string
}

// function to load configs
func LoadConfig() Config {
	cfg := Config{
		ServerPort: 3000,
		Environment: "development",
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	if app_env, exists := os.LookupEnv("APP_ENV"); exists {
		cfg.Environment = app_env
	}

	return cfg
}