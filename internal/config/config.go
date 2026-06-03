package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Addr string `yaml:"addr"`
	Mode string `yaml:"mode"`
}

type Data struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

type JWT struct {
	Secret string        `yaml:"secret"`
	Expire time.Duration `yaml:"expire"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}
	if cfg.Server.Addr == "" {
		cfg.Server.Addr = ":8080"
	}
	if cfg.Server.Mode == "" {
		cfg.Server.Mode = "debug"
	}
	if cfg.JWT.Expire == 0 {
		cfg.JWT.Expire = 24 * time.Hour
	}
	if cfg.Data.Driver == "" {
		cfg.Data.Driver = "mysql"
	}
	return &cfg, nil
}

func New(path string) (*Config, error) {
	return Load(path)
}
