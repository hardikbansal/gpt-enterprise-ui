package handlers

import (
	"encoding/json"
	"fmt"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
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
	logger.LogMessage(fmt.Sprintf("User Id : %s", userId))
	conversations, err := handler.srv.GetConversationsForUser(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, conversations)

}

func (handler *ApiHandler) CreateNewConversation(c *gin.Context) {
	userId := c.GetInt("user_id")
	// get POST request parameter from context c
	var data struct {
		Name string `json:"name"`
	}
	err := c.BindJSON(&data)
	logger.LogMessage(fmt.Sprintf("Conversation Name %s", data.Name))
	// Create New Conversation
	conversations, err := handler.srv.CreateNewConversation(userId, data.Name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, conversations)
}

func (handler *ApiHandler) DoQuery(c *gin.Context) {
	// get POST request parameter from context c
	var data struct {
		Query          string `json:"query"`
		ConversationId int    `json:"conversation_id"`
		Context        bool   `json:"context"`
	}
	err := c.BindJSON(&data)
	logger.LogMessage(fmt.Sprintf("data: %s", data))
	// Create New Conversation
	queries, err := handler.srv.QueryLLM(data.Query, data.ConversationId, data.Context)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, queries)
}

func (handler *ApiHandler) GetQueries(c *gin.Context) {
	var conversationId string
	conversationId, isPresent := c.GetQuery("conversationId")
	if !isPresent {
		conversationId = ""
	}
	conversationIdInt, err := strconv.Atoi(conversationId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "conversation_id must be an integer"})
		return
	}
	// Create New Conversation
	queries, err := handler.srv.GetQueriesForConversation(conversationIdInt)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, queries)
}

func (handler *ApiHandler) GetTemplates(c *gin.Context) {
	userId := c.GetInt("user_id")
	logger.LogMessage(fmt.Sprintf("User Id : %s", userId))
	templates, err := handler.srv.GetTemplates(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, templates)
}

func (handler *ApiHandler) StoreTemplate(c *gin.Context) {
	userId := c.GetInt("user_id")
	// get POST request parameter from context c
	var data struct {
		TemplateName string   `json:"name"`
		Parts        []string `json:"parts"`
		Params       []string `json:"params"`
	}
	err := c.BindJSON(&data)
	logger.LogMessage(fmt.Sprintf("User Id : %s params %s parts %s with template name %s", userId, data.Params, data.Parts, data.TemplateName))
	//c.AbortWithStatus(http.StatusInternalServerError)
	//return
	templates, err := handler.srv.StoreTemplate(userId, data.TemplateName, data.Parts, data.Params)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, templates)
}
