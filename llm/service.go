package llm

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/hardikbansal/gpt-enterprise-ui/config"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"io/ioutil"
	"net/http"
)

const OPEN_API_URL = "https://api.openai.com/v1/chat/completions"

type Service struct {
}

func currentModel() string {
	return "gpt-3.5-turbo-0301"
}

func (s *Service) Query(prompts []core.LLMPrompt) ([]byte, error) {
	// create the request payload
	data := map[string]any{
		"model":    currentModel(),
		"messages": prompts,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	// create the HTTP client
	client := http.Client{}

	// create the request object
	req, err := http.NewRequest("POST", OPEN_API_URL, bytes.NewBuffer(payload))
	if err != nil {
		logger.LogError(err)
		return nil, err
	}

	// set the API key in the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.GetInstance().ApiKey))

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	println(string(body))
	return body, nil
}

func (s *Service) ResponseToLLMPrompt(response []byte) (core.LLMPrompt, error) {
	var llmResponse LLMResponse
	err := json.Unmarshal(response, &llmResponse)
	if err != nil {
		return core.LLMPrompt{}, err
	}
	numMessages := len(llmResponse.Choices)
	if numMessages == 0 {
		return core.LLMPrompt{Role: "assistant", Content: "# Error"}, nil
	}
	return core.LLMPrompt{Role: llmResponse.Choices[0].Message.Role, Content: llmResponse.Choices[0].Message.Content}, nil
}
