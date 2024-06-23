package config

import (
	errors "github.com/rotisserie/eris"
	"github.com/spf13/viper"
)

func ReadConfig() error {
	viper.AddConfigPath("./manifest/config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

type Config struct{}
