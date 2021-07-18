package app

import (
	"github.com/RamiAwar/go_demo_users_api/routers/stats"
	"github.com/RamiAwar/go_demo_users_api/routers/user"
)

func mapRoutes(){
	router.GET("/health", stats.HealthCheck)
	router.GET("/version", stats.Version)
	
	router.POST("/user", user.CreateUser)
	router.GET("/user/:user_id", user.GetUser)
	router.GET("/user/autocomplete", user.AutocompleteUser)
}