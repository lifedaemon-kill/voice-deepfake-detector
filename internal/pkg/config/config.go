package config

import (
	"github.com/cockroachdb/errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const Path = "config/config.yaml"

type Config struct {
	Bot BotConfig `yaml:"bot"`
	App AppConfig `yaml:"app"`
	Env string    `yaml:"env"`
}

type BotConfig struct {
	Token string `yaml:"token"`
}
type AppConfig struct {
	AiServiceHost string `yaml:"ai_service_host"`
}

func Load(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "no such file %s", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}
	return &cfg, nil
}
