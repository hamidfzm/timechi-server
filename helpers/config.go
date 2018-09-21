package helpers

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog"
	"os"
)

type config struct {
	Port   int
	Secret string
	Debug  bool
	DBName string `mapstructure:"db_name"`
}

var Config config
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetDefault("port", 8080)
	viper.SetDefault("debug", false)
	viper.SetDefault("secret", "secret")
	viper.SetDefault("db_name", "data/timechi.db")
	viper.BindEnv("port", "PORT")
}

func initConfig() {
	viper.SetConfigName("config")
	
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("/etc/timechi/")
		viper.AddConfigPath("./data")
	}
	
	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Msgf("Can't read config: %s", err)
	}
	viper.Unmarshal(&Config)
	
	if Config.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
