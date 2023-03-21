package db

import "github.com/hardikbansal/gpt-enterprise-ui/core"

func (user *User) ToDomainUser() core.User {
	return core.User{
		Email:   user.Email,
		Name:    user.Name,
		Picture: "",
		ID:      user.ID,
	}
}
