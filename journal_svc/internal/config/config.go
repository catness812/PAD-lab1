package config

import (
	"os"

	"github.com/gookit/slog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"host"`
	GrpcPort int    `yaml:"grpc_port"`
	Database Mongo  `yaml:"mongo"`
}

type Mongo struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

var Cfg Config

func LoadConfig() {
	var errOccurred bool

	data, err := os.ReadFile("./journal_svc/config.yml")
	if err != nil {
		errOccurred = true
	}

	if errOccurred {
		data, err = os.ReadFile("../journal_svc/config.yml")
		if err == nil {
			errOccurred = false
		}
	}

	if errOccurred {
		slog.Error("Failed to read configuration file: %v", err)
	}

	if err := yaml.Unmarshal(data, &Cfg); err != nil {
		slog.Error("Failed to unmarshal YAML data: %v", err)
	}
}
