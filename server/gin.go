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

	// serve the chat UI
	//router.Use(static.Serve("/", static.LocalFile("./ui/build", true)))
	//router.Static("/", "./ui/build")

	// define the endpoint for the chat API
	router.POST("/api/accesstoken", handler.GetAccessToken)
	//router.POST("/chat", handler.CallChatGptApi)

	//router.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	//})
	logger.LogMessage("starting http server")
	router.Run("0.0.0.0:" + port)
}
