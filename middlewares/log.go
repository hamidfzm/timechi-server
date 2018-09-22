package middlewares

import (
	"net/http"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("method", r.Method).Str("path", r.URL.Path).Msg("Request")
		handler.ServeHTTP(w, r)
	})
}
