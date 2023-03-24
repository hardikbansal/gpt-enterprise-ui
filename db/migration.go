package db

import logger "github.com/hardikbansal/gpt-enterprise-ui/logger"

func (adapter *DbAdapter) RunMigrations() {
	err := adapter.db.AutoMigrate(&User{})
	if err != nil {
		logger.LogProblem("Error auto migrating user table")
		return
	}
	//do auto migration for conversation table
	err = adapter.db.AutoMigrate(&Conversation{})
	if err != nil {
		logger.LogProblem("Error auto migrating conversation table")
		return
	}
	//do auto migration for query table
	err = adapter.db.AutoMigrate(&Query{})
	if err != nil {
		logger.LogProblem("Error auto migrating query table")
		return
	}

	//do auto migration for templates table
	err = adapter.db.AutoMigrate(&Template{})
	if err != nil {
		logger.LogProblem("Error auto migrating template table")
	}
}
