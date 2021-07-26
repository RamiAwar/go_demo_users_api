package stats

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Version(c *gin.Context) {
	c.String(http.StatusOK, "v0.0.1")
}
