package services

import (
	"net/http"
	"github.com/hamidfzm/timechi-server/errors"
	"encoding/json"
	"fmt"
	"github.com/hamidfzm/timechi-server/config"
)

type Handler struct {
	*config.Config
}

func (h *Handler) Abort(w http.ResponseWriter, statusCode int) {
	message := http.StatusText(statusCode)
	http.Error(w, message, statusCode)
	panic(errors.Abort{message})
}

func (h *Handler) DecodeJsonRequest(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func (h *Handler) EncodeJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	if h.Config.Debug {
		if j, err := json.MarshalIndent(data, "", "\t"); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			fmt.Fprintf(w, "%s", j)
			
			return nil
		} else {
			return err
		}
	} else {
		if j, err := json.Marshal(data); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(statusCode)
			fmt.Fprintf(w, "%s", j)
			
			return nil
		} else {
			return err
		}
	}
	
}
