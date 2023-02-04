package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT              string `mapstructure:"PORT"`
	DATABASE_SCHEME   string `mapstructure:"DB_TYPE"`
	DATABASE_NAME     string `mapstructure:"DB_NAME"`
	DATABASE_HOST     string `mapstructure:"DB_HOST"`
	DATABASE_PORT     int    `mapstructure:"DB_PORT"`
	DATABASE_USERNAME string `mapstructure:"DB_USERNAME"`
	DATABASE_PASSWORD string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig() (config EnvVars, err error) {

	env := os.Getenv("GO_ENV")
	if env == "production" {
		port, errd := strconv.Atoi(os.Getenv("DB_PORT"))
		if errd != nil {
			return
		}
		return EnvVars{
			PORT:              os.Getenv("PORT"),
			DATABASE_SCHEME:   os.Getenv("DB_TYPE"),
			DATABASE_NAME:     os.Getenv("DB_NAME"),
			DATABASE_HOST:     os.Getenv("DB_HOST"),
			DATABASE_USERNAME: os.Getenv("DB_USERNAME"),
			DATABASE_PASSWORD: os.Getenv("DB_PASSWORD"),
			DATABASE_PORT:     port,
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// Let's validate config here
	if config.PORT == "" {
		err = errors.New("PORT must be specified")
		return
	}

	return
}
