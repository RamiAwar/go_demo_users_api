package app

import (
	"github.com/RamiAwar/go_demo_users_api/routers/stats"
	"github.com/RamiAwar/go_demo_users_api/routers/users"
)

func mapRoutes(){
	router.GET("/health", stats.HealthCheck)
	router.GET("/version", stats.Version)
	
	router.POST("/user", users.CreateUser)
	router.GET("/user/:user_id", users.GetUser)
	router.GET("/user/autocomplete", users.AutocompleteUser)
}