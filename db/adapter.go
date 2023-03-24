package db

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InvalidData struct {
	error string
}

func (error InvalidData) Error() string {
	return error.error
}

type DbAdapter struct {
	db *gorm.DB
}

func GetDbAdapter(dbConnUrl string) (*DbAdapter, error) {
	dbConnString := fmt.Sprintf(dbConnUrl)
	conn, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		return &DbAdapter{}, err
	}
	return &DbAdapter{conn}, nil
}

func (adapter *DbAdapter) GetOrCreateUser(email string, name string) (core.User, error) {
	if email == "" {
		return core.User{}, InvalidData{"Email not present in user detail"}
	}
	user, err := adapter.getUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user, err = adapter.createUser(email, name)
		return user.ToCoreUser(), err
	}
	return user.ToCoreUser(), err
}

func (adapter *DbAdapter) createUser(email string, name string) (User, error) {
	user := User{Email: email, Name: name}
	result := adapter.db.Create(&user)
	return user, result.Error
}

func (adapter *DbAdapter) getUserByEmail(email string) (User, error) {
	var user User
	result := adapter.db.Where(&User{Email: email}).First(&user)
	return user, result.Error
}

func (adapter *DbAdapter) GetUserById(userId int) (core.User, error) {
	var user User
	result := adapter.db.Where(&User{ID: uint(userId)}).First(&user)
	return user.ToCoreUser(), result.Error
}

// GetConversationByUser function to get conversation based on user and return it after mapping it to core conversation
// using ToCoreConversation function and also fetch only last 20 conversations based on CreatedAt field
func (adapter *DbAdapter) GetConversationByUser(userId int) ([]core.Conversation, error) {
	var conversations []Conversation
	result := adapter.db.Where(&Conversation{UserID: uint(userId)}).Order("created_at DESC").Limit(20).Find(&conversations)
	// loop on conversation and map conversation to core.conversation using toCoreConversation
	return Map(conversations, func(c Conversation) core.Conversation {
		return c.ToCoreConversation()
	}), result.Error
}

// GetQueriesByConversation function to get query based on ConversationID and return list after mapping it to core message
// model using ToCoreQuery function. Only fetch last 20 messages based on CreatedAt field
func (adapter *DbAdapter) GetQueriesByConversation(conversationId int) ([]core.Query, error) {
	var query []Query
	result := adapter.db.Where(&Query{ConversationID: uint(conversationId)}).Order("created_at DESC").Limit(20).Find(&query)
	return Map(query, func(q Query) core.Query {
		return q.ToCoreQuery()
	}), result.Error
}

func (adapter *DbAdapter) CreateNewConversation(userId int, conversationName string) error {
	result := adapter.db.Create(&Conversation{UserID: uint(userId), Name: conversationName})
	return result.Error
}

func (adapter *DbAdapter) StoreQueryForConversation(conversationId int, query string, response []byte, context int) error {
	result := adapter.db.Create(&Query{ConversationID: uint(conversationId), Query: query, Response: response, Context: int32(context)})
	return result.Error
}

func (adapter *DbAdapter) GetContextForQuery(conversationId int, context int) ([]core.Query, error) {
	var queries []Query
	result := adapter.db.Where(&Query{ConversationID: uint(conversationId)}).Order("created_at DESC").Limit(context).Find(&queries)
	return Map(queries, func(q Query) core.Query {
		return q.ToCoreQuery()
	}), result.Error
}

func (adapter *DbAdapter) GetTemplatesByUserId(userId int) ([]core.Template, error) {
	var templates []Template
	result := adapter.db.Where(&Template{UserID: uint(userId)}).Order("created_at DESC").Limit(20).Find(&templates)
	return Map(templates, func(q Template) core.Template {
		return q.toCore()
	}), result.Error
}

func (adapter *DbAdapter) StoreTemplate(templateName string, userId int, parts []string, params []string) error {
	properties, err := json.Marshal(TemplateProperties{Parts: parts, Params: params})
	if err != nil {
		return err
	}
	result := adapter.db.Create(&Template{UserID: uint(userId), Name: templateName, Properties: properties})
	return result.Error
}
