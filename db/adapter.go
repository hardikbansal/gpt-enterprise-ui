package db

import (
	"errors"
	"fmt"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InvalidData struct {
	error string
}

func (error InvalidData) Error() string {
	return error.error
}

type DbAdapter struct {
	db *gorm.DB
}

func GetDbAdapter(dbConnUrl string) (*DbAdapter, error) {
	dbConnString := fmt.Sprintf(dbConnUrl)
	conn, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {
		return &DbAdapter{}, err
	}
	return &DbAdapter{conn}, nil
}

func (adapter *DbAdapter) GetOrCreateUser(email string, name string) (core.User, error) {
	if email == "" {
		return core.User{}, InvalidData{"Email not present in user detail"}
	}
	user, err := adapter.getUserByEmail(email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user, err = adapter.createUser(email, name)
		return user.ToDomainUser(), err
	}
	return user.ToDomainUser(), err
}

func (adapter *DbAdapter) createUser(email string, name string) (User, error) {
	user := User{Email: email, Name: name}
	result := adapter.db.Create(&user)
	return user, result.Error
}

func (adapter *DbAdapter) getUserByEmail(email string) (User, error) {
	var user User
	result := adapter.db.Where(&User{Email: email}).First(&user)
	return user, result.Error
}
