package config

import (
	"github.com/spf13/viper"
	"os"
)

type structure struct {
	App struct {
		HttpPort string `mapstructure:"httpPort"`
		GrpcPort string `mapstructure:"grpcPort"`
	} `mapstructure:"app"`
}

type config struct {
	HttpPort string
	GrpcPort string
}

var cfg config

func GetConfig() config {
	return cfg
}

func Load() error {
	path, _ := os.Getwd()

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config := structure{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	cfg.HttpPort = config.App.HttpPort
	cfg.GrpcPort = config.App.GrpcPort

	return nil

}
