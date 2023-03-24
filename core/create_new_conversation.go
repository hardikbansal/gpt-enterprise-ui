package core

func (srv ChatGptService) CreateNewConversation(userId int, conversationName string) ([]Conversation, error) {
	err := srv.dbService.CreateNewConversation(userId, conversationName)
	if err != nil {
		return nil, err
	}
	return srv.GetConversationsForUser(userId)
}
