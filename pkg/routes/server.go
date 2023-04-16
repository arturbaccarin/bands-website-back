package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// Config configures the gin router
type Config struct{}

// setupRouter will configurate gin server
func (Config) setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	return r
}

// Start will start the router server
func (c Config) Start() {
	router = c.setupRouter()
	_ = router.Run(":8080")
}
