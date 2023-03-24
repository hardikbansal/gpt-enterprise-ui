package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func (handler *ApiHandler) GetConversations(c *gin.Context) {
	userId := c.GetInt("user_id")
	conversations, err := handler.srv.GetConversationsForUser(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	conversationsJson, err := json.Marshal(conversations)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, conversationsJson)

}

func (handler *ApiHandler) CreateNewConversation(c *gin.Context) {
	userId := c.GetInt("user_id")
	// get POST request parameter from context c
	conversationName := c.PostForm("conversation_name")
	// Create New Conversation
	conversations, err := handler.srv.CreateNewConversation(userId, conversationName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	conversationsJson, err := json.Marshal(conversations)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, conversationsJson)
}

func (handler *ApiHandler) DoQuery(c *gin.Context) {
	// get POST request parameter from context c
	query := c.PostForm("query")
	converationId, err := strconv.Atoi(c.PostForm("conversation_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "conversation_id must be an integer"})
		return
	}
	isContext, err := strconv.ParseBool(c.PostForm("context"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "context must be a boolean"})
	}
	// Create New Conversation
	conversations, err := handler.srv.QueryLLM(query, converationId, isContext)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	conversationsJson, err := json.Marshal(conversations)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, conversationsJson)
}

func (handler *ApiHandler) GetQueries(c *gin.Context) {
	conversationId, err := strconv.Atoi(c.PostForm("conversation_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "conversation_id must be an integer"})
		return
	}
	// Create New Conversation
	queries, err := handler.srv.GetQueriesForConversation(conversationId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	conversationsJson, err := json.Marshal(queries)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, conversationsJson)
}
