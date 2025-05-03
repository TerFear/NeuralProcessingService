package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env   string `yaml:"env"`
	Token string `yaml:"token"`
	GRPC  `yaml:"grpc"`
}

type GRPC struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func ReadConfig() *Config {
	path := GetPath()

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	return &cfg
}

func GetPath() (path string) {
	path = os.Getenv("CONFIG_PATH")

	if path == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist")
	}

	return path
}
