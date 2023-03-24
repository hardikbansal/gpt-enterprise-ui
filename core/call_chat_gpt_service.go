package core

import (
	"encoding/json"
)

func (srv *ChatGptService) QueryLLM(query string, conversationId int, isContext bool) ([]Query, error) {
	var context []Query
	if isContext {
		queries, err := srv.dbService.GetContextForQuery(conversationId, 4)
		context = append(context, queries...)
		if err != nil {
			return nil, err
		}
	}

	// preparing prompt based on context
	prompt := make([]LLMPrompt, 0, len(context)*2)

	for _, e := range context {
		prompt = append(prompt, LLMPrompt{Role: "user", Content: "Say this is a test!"})
		prompt = append(prompt, LLMPrompt{Role: "bot", Content: "Say this is a test!"})
	}

	prompt = append(prompt, LLMPrompt{Role: "user", Content: "Say this is a test!"})
	resp, err := srv.llmService.Query(prompt)
	if err != nil {
		return nil, err
	}
	// parse the response
	var respData struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}
	err = json.Unmarshal(resp, &respData)
	if err != nil {
		return nil, err
	}

	// store response in db
	err = srv.dbService.StoreQueryForConversation(conversationId, query, resp, isContext)
	if err != nil {
		return nil, err
	}

	return srv.GetQueriesForConversation(conversationId)

}
