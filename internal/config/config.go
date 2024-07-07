package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type BOTConfig struct {
	TOKEN string `mapstructure:"token"`
}

type Config struct {
	BOT BOTConfig `mapstructure:"bot"`
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

	err = errors.Wrap(viper.Unmarshal(&GlobalConfig), "failed to unmarshal the config")
	if err != nil {
		return err
	}

	return nil
}
