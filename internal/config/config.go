package config

import (
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

type Config struct {
	BOT    BOTConfig    `mapstructure:"bot"`
	OpenAI OpenaiConfig `mapstructure:"openai"`
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

	err = errors.Wrap(viper.Unmarshal(&GlobalConfig), "failed to unmarshal the config")
	if err != nil {
		return err
	}

	return nil
}
