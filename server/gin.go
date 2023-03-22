package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikbansal/gpt-enterprise-ui/config"
	"github.com/hardikbansal/gpt-enterprise-ui/handlers"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
)

func StartServer(handler *handlers.ApiHandler, port string) {

	if !config.GetInstance().IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(logger.GetGinSupport())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// define the endpoint for the chat API
	router.POST("/api/accesstoken", handler.GetAccessToken)
	group := router.Group("/api/", handler.AuthMiddleware())
	group.GET("/user", handler.GetUserDetails)
	logger.LogMessage("starting http server")
	err := router.Run("0.0.0.0:" + port)
	if err != nil {
		return
	}
}
