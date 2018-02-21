package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/jsons"
	"github.com/hamidfzm/timechi-server/helpers"
	"github.com/hamidfzm/timechi-server/models"
)

type UserController struct{}

func (c UserController) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var req jsons.RegisterV1
	helpers.DecodeJson(r, &req)
	
	var user models.User
	req.To(&user)
	
	var resp jsons.PublicProfileV1
	resp.From(&user)
	
	helpers.EncodeJson(w, &resp)
}
