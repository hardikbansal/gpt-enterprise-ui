package core

type Service interface {
	CallChatGPTAPI(query string, conversation_id int, is_context bool) (string, error)
	GenerateAccessToken(token string) (AccessTokenData, error)
	//GetUserDetails(id int) (User, error)
	//GetSavedPrompts(id int) ([]string, error)
	//GetConversationMessages(id int)
}

type DbService interface {
	GetOrCreateUser(email string, name string) (User, error)
}
