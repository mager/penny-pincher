package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Auth0Domain   string
	Auth0Audience string
	Auth0Secret   string
	BitIOAPIKey   string
	NextSecret    string
}

func ProvideConfig(log *zap.SugaredLogger) Config {
	var cfg Config
	err := envconfig.Process("pennypincher", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}

var Options = ProvideConfig
