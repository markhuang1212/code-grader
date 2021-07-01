package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	authorized := r.Group("/")

	authorized.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authorized.POST("/grade", func(c *gin.Context) {

	})

	return r
}
