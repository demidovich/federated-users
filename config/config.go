package config

import (
	"federated/pkg/db/postgres"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres postgres.Config
}

func New(filename string) (Config, error) {
	log.Printf("Init app configuration from %s", filename)

	v := viper.New()
	v.SetConfigFile(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return Config{}, fmt.Errorf("unable to decode config into struct, %v", err)
	}

	return c, nil
}
