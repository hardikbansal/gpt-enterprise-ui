package core

func (s *ChatGptService) VerifyAccessToken(token string) (User, error) {
	user, err := s.authService.VerifyAuthToken(token)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
