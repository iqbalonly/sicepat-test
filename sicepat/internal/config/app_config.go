package config

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	configPath = "./config.json"
)

func NewConfig() *AppConfig {

	config := &AppConfig{
		environment: "test",
		shutdownCh:  make(chan struct{}, 1),
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("failed to load config")
		panic(err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

type AppConfig struct {
	environment string
	Server      string     `json:"server"`
	Mysql       *MysqlConf `json:"mysql"`

	shutdownCh chan struct{}
}

func (a *AppConfig) Shutdown(ctx context.Context) {
	close(a.shutdownCh)
}

func (a *AppConfig) Environment() string {
	return a.environment
}

func (a *AppConfig) ServerAddress() string {
	if a.Server == "" {
		return "0.0.0.0:9999"
	}

	return a.Server
}
