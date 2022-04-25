package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	ParseConfig,
)

func ParseConfig(filePath string) {
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}
