package core

func (s *ChatGptService) GetTemplates(userId int) ([]Template, error) {
	templates, err := s.dbService.GetTemplatesByUserId(userId)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (s *ChatGptService) StoreTemplate(userId int, templateName string, parts []string, params []string) ([]Template, error) {
	err := s.dbService.StoreTemplate(templateName, userId, parts, params)
	if err != nil {
		return nil, err
	}
	return s.GetTemplates(userId)
}
