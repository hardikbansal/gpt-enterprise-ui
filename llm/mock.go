package llm

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
)

type MockService struct {
}

func (s *MockService) Query(prompts []core.LLMPrompt) ([]byte, error) {
	resp := []byte("{\n    \"id\": \"chatcmpl-6xYxmOsg3I47PSjQFZjB2hts3WLeV\",\n    \"object\": \"chat.completion\",\n    \"created\": 1679653770,\n    \"model\": \"gpt-3.5-turbo-0301\",\n    \"usage\": {\n        \"prompt_tokens\": 14,\n        \"completion_tokens\": 5,\n        \"total_tokens\": 19\n    },\n    \"choices\": [\n        {\n            \"message\": {\n                \"role\": \"assistant\",\n                \"content\": \"This is a test!\"\n            },\n            \"finish_reason\": \"stop\",\n            \"index\": 0\n        }\n    ]\n}")
	return resp, nil
}

func (s *MockService) ResponseToLLMPrompt(response []byte) (core.LLMPrompt, error) {
	var llmResponse LLMResponse
	logger.LogMessage(fmt.Sprintf("%s", string(response)))
	err := json.Unmarshal(response, &llmResponse)
	if err != nil {
		logger.LogMessage(fmt.Sprintf("%s", err))
		return core.LLMPrompt{}, err
	}
	numMessages := len(llmResponse.Choices)
	if numMessages == 0 {
		return core.LLMPrompt{Role: "bot", Content: ""}, nil
	}
	logger.LogMessage(fmt.Sprintf("%s", llmResponse.Choices))
	return core.LLMPrompt{Role: llmResponse.Choices[0].Message.Role, Content: llmResponse.Choices[0].Message.Content}, nil
}
