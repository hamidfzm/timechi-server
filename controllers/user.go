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

func (c UserControllerV1) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req jsons.LoginV1
	
	if err := helpers.DecodeJsonRequest(r, &req); err != nil {
		helpers.Abort(w, http.StatusBadRequest)
	}
	
	if err := helpers.Validate.Struct(req); err != nil {
		helpers.Abort(w, http.StatusUnprocessableEntity)
	}
	
	if user, err := models.FindUserByEmail(req.Email); err != nil {
		helpers.Abort(w, http.StatusNotFound)
	} else if err := user.VerifyPassword(req.Password); err == nil {
		token, _ := user.GenerateToken()
		// TODO handle generate token error
		helpers.EncodeJsonResponse(w, http.StatusOK, map[string]string{"token": token})
		
	} else {
		helpers.Abort(w, http.StatusUnauthorized)
		
	}
}

func (c UserControllerV1) Profile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp jsons.PublicProfileV1
	
	user, _ := r.Context().Value("user").(*models.User)
	resp.From(user)
	
	helpers.EncodeJsonResponse(w, http.StatusOK, resp)
}
