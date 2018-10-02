package middlewares

import (
	"net/http"
	"github.com/rs/zerolog/log"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, 200}
		
		start := time.Now()
		handler.ServeHTTP(&rec, r)
		end := time.Now()
		
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("code", rec.status).
			TimeDiff("duration", end, start).
			Msg("Request")
	})
}
