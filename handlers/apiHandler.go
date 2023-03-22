package handlers

import (
	"encoding/json"
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
				return
			}
		default:
			{
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, jwtToken)

}

func (handler *ApiHandler) GetUserDetails(c *gin.Context) {
	userId := c.GetInt("user_id")
	user, err := handler.srv.GetUserDetails(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userData, err := json.Marshal(user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, userData)
}
