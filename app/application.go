package app

import (
	"github.com/RamiAwar/go_demo_users_api/utils/logging"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Start() {
	mapRoutes()
	logging.Info("Starting application...")
	router.Run(":9000")
}
