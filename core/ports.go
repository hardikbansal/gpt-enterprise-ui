package core

type Service interface {
	QueryLLM(query string, conversationId int, isContext bool) ([]Query, error)
	GenerateAccessToken(token string) (AccessTokenData, error)
	VerifyAccessToken(token string) (User, error)
	GetUserDetails(userId int) (User, error)
	GetConversationsForUser(userId int) ([]Conversation, error)
	GetQueriesForConversation(conversationId int) ([]Query, error)
	CreateNewConversation(userId int, conversationName string) ([]Conversation, error)
	GetTemplates(userId int) ([]Template, error)
	StoreTemplate(userId int, templateName string, parts []string, params []string) ([]Template, error)
}

type DbService interface {
	GetOrCreateUser(email string, name string) (User, error)
	GetUserById(userId int) (User, error)
	GetConversationByUser(userId int) ([]Conversation, error)
	GetQueriesByConversation(conversationId int) ([]Query, error)
	CreateNewConversation(userId int, conversationName string) error
	StoreQueryForConversation(conversationID int, query string, response []byte, context int) error
	GetContextForQuery(conversationId int, maxContext int) ([]Query, error)
	GetTemplatesByUserId(userId int) ([]Template, error)
	StoreTemplate(templateName string, userId int, parts []string, params []string) error
}

type AuthService interface {
	GenerateAuthToken(user User, expiry int64) (string, error)
	VerifyAuthToken(token string) (User, error)
}

type LLMService interface {
	Query(prompt []LLMPrompt) ([]byte, error)
	ResponseToLLMPrompt(response []byte) (LLMPrompt, error)
}
