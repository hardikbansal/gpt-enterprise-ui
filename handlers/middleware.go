package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"net/http"
	"strings"
)

func splitBearerToken(bearerToken string) (string, error) {
	// Split the bearer token string into two parts
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		// The bearer token should be in the format "Bearer <token>"
		return "", fmt.Errorf("Invalid bearer token")
	}
	return parts[1], nil
}

func (handler *ApiHandler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get the Bearer token from the header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		logger.LogMessage(fmt.Sprintf("Auth Header %s", authHeader))
		token, err := splitBearerToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		user, err := handler.srv.VerifyAccessToken(token)
		if _, ok := err.(core.InvalidAuthToken); ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		logger.LogMessage(fmt.Sprintf("User %s", user))
		c.Set("user_id", int(user.ID))
		c.Next()
	}
}
