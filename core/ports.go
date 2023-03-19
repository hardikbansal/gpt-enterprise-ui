package core

type Service interface {
	CallChatGPTAPI(query string) (string, error)
	GenerateAccessToken(token string) (AccessTokendata, error)
}
