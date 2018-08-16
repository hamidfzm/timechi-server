package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/jsons"
)

type TimeController struct{}

func (c TimeController) StartTimer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req jsons.StartTimerV1
	
	if err := helpers.DecodeJsonRequest(r, &req); err != nil {
		helpers.Abort(w, http.StatusBadRequest)
	}
	
	if err := helpers.Validate.Struct(req); err != nil {
		helpers.Abort(w, http.StatusUnprocessableEntity)
	}
	
	user := currentUser(r)
	if err := user.StartTimer(req.Title); err != nil {
		helpers.Abort(w, http.StatusConflict)
	}
	
	var resp jsons.TimeV1
	resp.Title = *user.TimerTitle
	resp.StartedAt = helpers.JSONTime{Time: *user.TimerStartAt}
	
	helpers.EncodeJsonResponse(w, http.StatusOK, resp)
}

func (c TimeController) StopTimer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	user := currentUser(r)
	if time, err := user.StopTimer(); err != nil {
		helpers.Abort(w, http.StatusConflict)
	} else {
		var resp jsons.TimeV1
		
		resp.From(&time)
		helpers.EncodeJsonResponse(w, http.StatusOK, resp)
	}
}

func (c TimeController) Times(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp jsons.TimesV1
	
	user := currentUser(r)
	times := user.Times()
	
	// TODO implement pagination
	
	resp.From(&times, 1, 10, 10)
	helpers.EncodeJsonResponse(w, http.StatusOK, resp)
}
