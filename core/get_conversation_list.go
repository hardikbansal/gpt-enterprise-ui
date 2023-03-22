package core

// Generate GetConversationList function. It is a function over ChatGptStruct
// It takes userid as the parameter and fetches conversation using
// GetConversationByUser function of dbService present in struct
func (c *ChatGptService) GenerateGetConversationList(userid int) ([]Conversation, error) {
	conversations, err := c.dbService.GetConversationByUser(userid)
	if err != nil {
		return nil, err
	}
	return conversations, nil
}
