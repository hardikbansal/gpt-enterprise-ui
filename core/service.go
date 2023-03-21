package core

type UnauthorizedError struct {
}

func (e UnauthorizedError) Error() string {
	return "UnAuthorized"
}

type ChatGptService struct {
	dbService DbService
}

func GetNewService(dbService DbService) *ChatGptService {
	return &ChatGptService{dbService: dbService}
}
