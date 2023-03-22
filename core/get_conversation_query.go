package core

// function on ChatGptService to get queries based on conversation id
func (srv *ChatGptService) GetConversationQueries(conversationId int) ([]Query, error) {
	// calls GetQueriesByConversation function of dbservice and returns it
	queries, err := srv.dbService.GetQueriesByConversation(conversationId)
	if err != nil {
		return nil, err
	}
	return queries, nil
}
