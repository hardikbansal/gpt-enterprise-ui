package core

type UnauthorizedError struct {
}

func (e UnauthorizedError) Error() string {
	return "UnAuthorized"
}

type ChatGptService struct {
	dbService   DbService
	authService AuthService
}

func GetNewService(dbService DbService, authService AuthService) *ChatGptService {
	return &ChatGptService{dbService: dbService, authService: authService}
}
