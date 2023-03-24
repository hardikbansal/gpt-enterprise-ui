package core

import (
	"encoding/json"
)

func reverse(arr []Query) {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func (srv *ChatGptService) QueryLLM(query string, conversationId int, isContext bool) ([]Query, error) {
	var context []Query
	var numContext = 0
	if isContext {
		numContext = 4
		queries, err := srv.dbService.GetContextForQuery(conversationId, 4)
		context = append(context, queries...)
		if err != nil {
			return nil, err
		}
	}

	// preparing prompt based on context
	prompt := make([]LLMPrompt, 0, len(context)*2)
	reverse(context)
	for _, c := range context {
		prompt = append(prompt, LLMPrompt{Role: "user", Content: c.Query})
		resp, err := srv.llmService.ResponseToLLMPrompt([]byte(c.Response))
		if err != nil {
			prompt = append(prompt, LLMPrompt{Role: "assistant", Content: "# ERROR"})
		}
		prompt = append(prompt, resp)
	}

	prompt = append(prompt, LLMPrompt{Role: "user", Content: query})
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
	err = srv.dbService.StoreQueryForConversation(conversationId, query, resp, numContext)
	if err != nil {
		return nil, err
	}

	return srv.GetQueriesForConversation(conversationId)

}
