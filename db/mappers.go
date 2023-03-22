package db

import "github.com/hardikbansal/gpt-enterprise-ui/core"

func (user *User) ToCoreUser() core.User {
	return core.User{
		Email:   user.Email,
		Name:    user.Name,
		Picture: "",
		ID:      user.ID,
	}
}

// ToCoreConversation mapper to convert db model conversation to core.Conversation
func (conversation *Conversation) ToCoreConversation() core.Conversation {
	return core.Conversation{
		ID:        conversation.ID,
		Name:      conversation.Name,
		CreatedAt: conversation.CreatedAt.String(),
	}
}

// ToCoreQuery mapper to convert db model conversation to core.Query
func (query *Query) ToCoreQuery() core.Query {
	return core.Query{
		ID:        query.ID,
		Query:     query.Query,
		Response:  string(query.Response),
		CreatedAt: query.CreatedAt.String(),
	}
}
