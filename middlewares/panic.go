package middlewares

import (
	"net/http"
	"github.com/hamidfzm/timechi-server/helpers"
	"fmt"
)

func PanicMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				switch err.(type) {
				case helpers.ErrorAbort:
					return
				default:
					if helpers.Config.Debug {
						panic(err)
					}
					http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				}
			}()
			handler.ServeHTTP(w, r)
		})
}
