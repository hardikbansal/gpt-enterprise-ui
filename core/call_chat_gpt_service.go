package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (srv *ChatGptService) CallChatGPTAPI(query string, conversation_id int, is_context bool) (string, error) {
	// create the request payload
	data := map[string]string{
		"prompt": query,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// create the HTTP client
	client := http.Client{}

	// create the request object
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/davinci-codex/completions", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	// set the API key in the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_API_KEY_HERE")

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// parse the response
	var respData struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return "", err
	}

	// return the response text
	if len(respData.Choices) > 0 {
		return respData.Choices[0].Text, nil
	} else {
		return "", nil
	}
}
