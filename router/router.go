package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/controllers"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()
	
	userController := controllers.UserControllerV1{}
	Router.POST("/v1/user", userController.Register)
}
