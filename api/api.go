package api

import (
	"github.com/gin-gonic/gin"
)

func Route(rGroup *gin.RouterGroup) {

	//health check
	rGroup.GET("/health", healthHandler)
}

func healthHandler(c *gin.Context) {
	c.String(200, "Success")
}
