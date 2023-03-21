package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
)

type ApiHandler struct {
	srv core.Service
}

func NewApiHandler(service core.Service) *ApiHandler {
	return &ApiHandler{
		srv: service,
	}
}

func (handler *ApiHandler) CallChatGptApi(c *gin.Context) {
	// parse the request body
	var req core.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call the OpenAI API
	resp, err := handler.srv.CallChatGPTAPI(req.Query, 1, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// build the response
	res := core.ChatResponse{Response: resp}

	// send the response
	c.JSON(http.StatusOK, res)
}

func (handler *ApiHandler) GetAccessToken(c *gin.Context) {
	// Get the Bearer token from the header
	authHeader := c.GetHeader("Authorization")
	fmt.Println(authHeader)
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	tokenString := authHeader[len("Bearer "):]

	jwtToken, err := handler.srv.GenerateAccessToken(tokenString)
	if err != nil {
		switch err.(type) {
		case core.UnauthorizedError:
			{
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			}
		default:
			{
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			}
		}
	}
	c.JSON(http.StatusOK, jwtToken)

}
