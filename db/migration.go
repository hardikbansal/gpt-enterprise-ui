package db

import (
	"fmt"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations() {
	dbConnString := fmt.Sprintf("postgres://chatgpt:chatgpt@localhost:5432/chatgpt")
	conn, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if err != nil {

	}
	err = conn.AutoMigrate(&User{})
	if err != nil {
		logger.LogProblem("Error auto migrating user table")
		return
	}
	//conn.AutoMigrate(&User{})
	//conn.AutoMigrate(&User{})
	//conn.AutoMigrate(&User{})
}
