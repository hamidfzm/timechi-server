package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/helpers"
)

var Router *httprouter.Router

func SetupRouter() {
	Router = httprouter.New()
	Router.PanicHandler = helpers.PanicHandler
	
	userController := UserControllerV1{}
	Router.POST("/v1/user/register", userController.Register)
	Router.POST("/v1/user/login", userController.Login)
	Router.GET("/v1/user", authenticate(userController.Profile))
	
	timeController := TimeController{}
	Router.POST("/v1/time", authenticate(timeController.StartTimer))
	Router.DELETE("/v1/time", authenticate(timeController.StopTimer))
	Router.GET("/v1/time", authenticate(timeController.Times))
}
