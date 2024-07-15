package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type BOTConfig struct {
	TOKEN string `mapstructure:"token"`
}

type OpenaiConfig struct {
	APIKey string `mapstructure:"apikey"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	NAME     string `mapstructure:"database_name"`
	USERNAME string `mapstructure:"username"`
	PASSWORD string `mapstructure:"password"`
}

type Config struct {
	BOT    BOTConfig    `mapstructure:"bot"`
	OpenAI OpenaiConfig `mapstructure:"openai"`
	Mongo  MongoConfig  `mapstructure:"mongo"`
}

var GlobalConfig *Config

func Load() error {
	GlobalConfig = &Config{}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := errors.Wrap(
		viper.BindEnv("bot.token", "BOT.TOKEN"), "failed to bind BOT.TOKEN env")
	if err != nil {
		return err
	}

	err = errors.Wrap(
		viper.BindEnv("openai.apikey", "OPENAI_APIKEY"), "failed to bind OPENAI.APIKEY env")
	if err != nil {
		return err
	}

	err = errors.Wrap(
		viper.BindEnv("mongo.uri", "MONGO_URI"), "failed to bind MONGO_URI env")
	if err != nil {
		return err
	}

	err = errors.Wrap(
		viper.BindEnv("mongo.database_name", "MONGO_DATABASE_NAME"), "failed to bind MONGO_DATABASE_NAME env")
	if err != nil {
		return err
	}

	err = errors.Wrap(
		viper.BindEnv("mongo.username", "MONGO_USERNAME"), "failed to bind MONGO_USERNAME env")
	if err != nil {
		return err
	}

	err = errors.Wrap(
		viper.BindEnv("mongo.password", "MONGO_PASSWORD"), "failed to bind MONGO_PASSWORD env")
	if err != nil {
		return err
	}

	err = errors.Wrap(viper.Unmarshal(&GlobalConfig), "failed to unmarshal the config")
	if err != nil {
		return err
	}

	fmt.Println(GlobalConfig.Mongo.NAME)

	return nil
}
