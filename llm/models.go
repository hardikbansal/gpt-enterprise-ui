package llm

type JSONUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type LLMResponse struct {
	Id      string    `json:"id"`
	Model   string    `json:"model"`
	Usage   JSONUsage `json:"usage"`
	Choices []Choice  `json:"choices"`
}
