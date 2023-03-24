package db

import (
	"github.com/goccy/go-json"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
)

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
		ID:             query.ID,
		Query:          query.Query,
		ConversationId: query.ConversationID,
		Response:       string(query.Response),
		CreatedAt:      query.CreatedAt.String(),
	}
}

func (t *Template) toCore() core.Template {
	var props TemplateProperties
	err := json.Unmarshal(t.Properties, &props)
	if err != nil {
		return core.Template{}
	}
	return core.Template{
		ID:     t.ID,
		Name:   t.Name,
		Params: props.Params,
		Parts:  props.Parts,
	}
}
