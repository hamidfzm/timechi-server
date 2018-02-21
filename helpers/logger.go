package helpers

import (
	"net/http"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware(handler http.Handler) http.Handler {
	ourFunc := func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("method", r.Method).Str("path", r.URL.Path).Msg("Request")
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(ourFunc)
}
