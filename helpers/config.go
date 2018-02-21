package helpers

import "github.com/spf13/viper"

type config struct {
	Port      int
	PrettyLog bool `mapstructure:"pretty_log"`
}

var Config *config

func init() {
	viper.Set("port", 8080)
	viper.Set("pretty_log", true)
	viper.Unmarshal(&Config)
}
