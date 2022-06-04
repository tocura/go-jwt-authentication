package env

import (
	"context"
	"io/ioutil"

	"github.com/tocura/go-jwt-authentication/pkg/log"
	"gopkg.in/yaml.v2"
)

const configFileName = "config.yaml"

type App struct {
	Scope       string `yaml:"scope"`
	HTTPAddress string `yaml:"address"`
}

type Database struct {
	MySQL MySQLConfig `yaml:"mysql"`
}

type MySQLConfig struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
}

func New() (*Config, error) {
	var cfg Config

	src, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Error(context.TODO(), "error to read config file", err)
		return nil, err
	}

	if err := yaml.Unmarshal(src, &cfg); err != nil {
		log.Error(context.TODO(), "error to bind configs", err)
		return nil, err
	}

	log.Info(context.TODO(), "application configs initiated with success")
	return &cfg, nil
}
