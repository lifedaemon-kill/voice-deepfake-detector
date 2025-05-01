package config

import (
	"github.com/cockroachdb/errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

const ConfigPath = "../config/config.yaml"

type Config struct {
	Bot BotConfig `yaml:"bot"`
	App AppConfig `yaml:"app"`
	Env string    `yaml:"env"`
}

type BotConfig struct {
	Token string
}
type AppConfig struct {
	ASHost       string `yaml:"anti_spoofing_host"`
	AASEndpoint  string `yaml:"anti_audio_spoofing_endpoint"`
	AudioTempDir string `yaml:"temp_audio_dir"`
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

	cfg.Bot.Token = os.Getenv("BOT_TOKEN")
	return &cfg, nil
}
