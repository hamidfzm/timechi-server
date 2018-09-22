package config

import (
	"github.com/spf13/viper"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"os"
)

type Config struct {
	Host   string
	Port   int
	Secret string
	Debug  bool
	DBName string `mapstructure:"db_name"`
}

func NewConfig() *Config {
	viper.SetDefault(Host, "127.0.0.1")
	viper.SetDefault(Port, 8080)
	viper.SetDefault(Debug, false)
	viper.SetDefault("secret", "secret")
	viper.SetDefault("db_name", "data/timechi.db")
	viper.BindEnv(Host)
	viper.BindEnv(Port)
	
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	
	viper.ReadInConfig()
	config := Config{}
	viper.Unmarshal(&config)
	
	if config.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	return &config
}
