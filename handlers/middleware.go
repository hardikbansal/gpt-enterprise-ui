package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	"net/http"
)

func (handler *ApiHandler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get the Bearer token from the header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		user, err := handler.srv.VerifyAccessToken(authHeader)
		if _, ok := err.(core.InvalidAuthToken); ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
		c.Set("user", user)
		c.Next()
	}
}
