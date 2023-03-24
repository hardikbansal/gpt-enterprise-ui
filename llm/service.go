package llm

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"io/ioutil"
	"net/http"
)

const OPEN_API_URL = "https://api.openai.com/v1/chat/completions"
const API_KEY = "sk-temp"

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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))

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
	println("%s", string(body))
	return body, nil
}
