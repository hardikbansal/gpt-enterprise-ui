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

type AccessTokendata struct {
	Token string `json:"token"`
}

func GetEmptyAccessToken() AccessTokendata {
	return AccessTokendata{Token: ""}
}
