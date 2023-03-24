package core

import (
	"fmt"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
)

func (srv *ChatGptService) GetQueriesForConversation(conversationId int) ([]Query, error) {
	queries, err := srv.dbService.GetQueriesByConversation(conversationId)
	newQueries := Map(queries, func(query Query) Query {
		prompt, _ := srv.llmService.ResponseToLLMPrompt([]byte(query.Response))
		query.Response = prompt.Content
		logger.LogMessage(fmt.Sprintf("%s", query.Response))
		return query
	})
	if err != nil {
		return nil, err
	}
	return newQueries, nil
}
