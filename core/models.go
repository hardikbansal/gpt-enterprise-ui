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

type Conversation struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type Query struct {
	ID             uint   `json:"id"`
	ConversationId uint   `json:"conversation_id"`
	Query          string `json:"query"`
	Response       string `json:"response"`
	IsContext      string `json:"is_context"`
	CreatedAt      string `json:"created_at"`
}

type LLMPrompt struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AccessTokenData struct {
	Token string `json:"token"`
}

type LLMResponse struct {
	Response  []LLMPrompt `json:"response"`
	Id        string      `json:"id"`
	CreatedAt string      `json:"created_at"`
	Model     string      `json:"model"`
	Usage     string      `json:"usage"`
}
