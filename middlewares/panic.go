package middlewares

import (
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"github.com/hamidfzm/timechi-server/config"
	"github.com/hamidfzm/timechi-server/errors"
)

func PanicMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					switch err.(type) {
					case errors.Abort:
						return
					default:
						if viper.GetBool(config.Debug) {
							panic(err)
						}
						http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
					}
				}
			}()
			handler.ServeHTTP(w, r)
		})
}
