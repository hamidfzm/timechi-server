package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"github.com/hamidfzm/timechi-server/router"
	"github.com/hamidfzm/timechi-server/helpers"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	webCmd.Flags().IntP("port", "p", 0, "Web server port number")
	viper.BindPFlag("port", webCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(webCmd)
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start web server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf(":%d", helpers.Config.Port)
		log.Info().Msgf("Listening to %s", addr)
		if err := http.ListenAndServe(
			addr,
			buildChain(
				router.Router,
				helpers.LoggerMiddleware,
			)); err != nil {
			log.Fatal().Msgf("Listening failed: %s", err)
		}
	},
}

func buildChain(r http.Handler, ms ...func(handler http.Handler) http.Handler) http.Handler {
	if len(ms) == 0 {
		return r
	}
	return ms[0](buildChain(r, ms[1:cap(ms)]...))
	
}
