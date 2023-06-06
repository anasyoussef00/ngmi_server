package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	serverPort = 8080
)

type Config struct {
	ServerPort   int    `yaml:"server_port" env:"SERVER_PORT"`
	DSN          string `yaml:"dsn" env:"DSN,secret"`
	JwtSecretKey string `yaml:"jwt_secret_key" env:"JWT_SECRET_KEY,secret"`
}

func Load(configPath string) (*Config, error) {
	c := Config{ServerPort: serverPort}

	if content, err := os.ReadFile(configPath); err != nil {
		return nil, err
	} else {
		if err = yaml.Unmarshal(content, &c); err != nil {
			return nil, err
		}
		return &c, err
	}

}
