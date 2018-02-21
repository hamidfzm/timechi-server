package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/controllers"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()
	
	userController := controllers.UserController{}
	Router.POST("/user", userController.Register)
}
