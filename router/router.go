package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/hamidfzm/timechi-server/controllers"
	"github.com/hamidfzm/timechi-server/helpers"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()
	Router.PanicHandler = helpers.PanicHandler
	
	userController := controllers.UserControllerV1{}
	Router.POST("/v1/user", userController.Register)
}
