package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/jsons"
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/models"
)

type UserControllerV1 struct{}

func (c UserControllerV1) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req jsons.RegisterV1
	
	if err := helpers.DecodeJsonRequest(r, &req); err != nil {
		helpers.Abort(w, http.StatusBadRequest)
	}
	
	if err := helpers.Validate.Struct(req); err != nil {
		helpers.Abort(w, http.StatusUnprocessableEntity)
	}
	
	var user models.User
	req.To(&user)
	
	if err := user.Create(); err != nil {
		helpers.Abort(w, http.StatusConflict)
	}
	
	var resp jsons.PublicProfileV1
	resp.From(&user)
	
	helpers.EncodeJsonResponse(w, http.StatusCreated, &resp)
}
