package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/jsons"
)

type TimeController struct{}

func (c TimeController) StartTimer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp jsons.TimeV1
	
	user := currentUser(r)
	if err := user.StartTimer("test"); err != nil {
		helpers.Abort(w, http.StatusConflict)
	}
	
	//resp.From()
	helpers.EncodeJsonResponse(w, http.StatusOK, resp)
}

func (c TimeController) StopTimer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp jsons.TimeV1
	
	user := currentUser(r)
	if err := user.StopTimer(); err != nil {
		helpers.Abort(w, http.StatusConflict)
	}
	
	//resp.From()
	helpers.EncodeJsonResponse(w, http.StatusOK, resp)
}

func (c TimeController) Times(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp jsons.TimesV1
	
	user := currentUser(r)
	times := user.Times()
	resp.From(&times, 1, 10, 10)
	helpers.EncodeJsonResponse(w, http.StatusOK, resp)
}
