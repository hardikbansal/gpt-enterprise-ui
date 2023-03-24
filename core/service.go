package core

type UnauthorizedError struct {
}

func (e UnauthorizedError) Error() string {
	return "UnAuthorized"
}

type ChatGptService struct {
	dbService   DbService
	authService AuthService
	llmService  LLMService
}

func GetNewService(dbService DbService, authService AuthService, llmService LLMService) *ChatGptService {
	return &ChatGptService{dbService: dbService, authService: authService, llmService: llmService}
}
