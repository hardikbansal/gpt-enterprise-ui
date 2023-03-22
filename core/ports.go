package core

type Service interface {
	CallChatGPTAPI(query string, conversation_id int, is_context bool) (string, error)
	GenerateAccessToken(token string) (AccessTokenData, error)
	VerifyAccessToken(token string) (User, error)
	GetUserDetails(userId int) (User, error)
	GetConversationList(user User) ([]Conversation, error)
	//GetSavedPrompts(id int) ([]string, error)
	GetConversationQueries(conversationId int) ([]Query, error)
}

type DbService interface {
	GetOrCreateUser(email string, name string) (User, error)
	GetUserById(userId int) (User, error)
	GetConversationByUser(userId int) ([]Conversation, error)
	GetQueriesByConversation(conversationId int) ([]Query, error)
}

type AuthService interface {
	GenerateAuthToken(user User, expiry int64) (string, error)
	VerifyAuthToken(token string) (User, error)
}
