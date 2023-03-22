package core

func (s *ChatGptService) GetUserDetails(userId int) (User, error) {
	user, err := s.dbService.GetUserById(userId)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
