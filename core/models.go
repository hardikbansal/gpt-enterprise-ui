package core

// define the request
type ChatRequest struct {
	Query string `json:"query"`
}

// define the response structure
type ChatResponse struct {
	Response string `json:"response"`
}

type GoogleUserInfo struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type User struct {
	ID      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type AccessTokenData struct {
	Token string `json:"token"`
}
