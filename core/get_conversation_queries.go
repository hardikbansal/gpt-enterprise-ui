package core

func (srv *ChatGptService) GetQueriesForConversation(conversationId int) ([]Query, error) {
	queries, err := srv.dbService.GetQueriesByConversation(conversationId)
	if err != nil {
		return nil, err
	}
	return queries, nil
}
