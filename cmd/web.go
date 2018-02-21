package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
	"github.com/hamidfzm/timechi-server/router"
	"github.com/hamidfzm/timechi-server/helpers"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"context"
	"time"
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
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		
		addr := fmt.Sprintf(":%d", helpers.Config.Port)
		
		log.Info().Msgf("Listening on %s", addr)
		
		server := &http.Server{Addr: addr, Handler: buildChain(
			router.Router,
			helpers.LoggerMiddleware,
		)}
		
		go func() {
			if err := server.ListenAndServe(); err != nil {
				log.Fatal().Msgf("%s", err)
			}
		}()
		<-stop
		
		log.Info().Msg("Shutting down the server...")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		server.Shutdown(ctx)
		log.Info().Msg("Server gracefully stopped")
	},
}

func buildChain(r http.Handler, ms ...func(handler http.Handler) http.Handler) http.Handler {
	if len(ms) == 0 {
		return r
	}
	return ms[0](buildChain(r, ms[1:cap(ms)]...))
	
}
