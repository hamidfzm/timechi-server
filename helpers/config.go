package helpers

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"os"
)

type config struct {
	Port      int
	PrettyLog bool `mapstructure:"pretty_log"`
	Secret    string
}

var Config config
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault("port", 8080)
	viper.SetDefault("pretty_log", true)
	viper.SetDefault("secret", "secret")
}

func initConfig() {
	viper.SetConfigName("config")
	
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("/etc/timechi/")
		viper.AddConfigPath(".")
	}
	
	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Msgf("Can't read config: %s", err)
	}
	viper.Unmarshal(&Config)
	
	if Config.PrettyLog {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
