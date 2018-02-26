package helpers

import (
	"net/http"
)

func PanicHandler(w http.ResponseWriter, r *http.Request, error interface{}) {
	switch error.(type) {
	case ErrorAbort:
		return
	default:
		panic(error)
	}
	
}
