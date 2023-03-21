package db

import logger "github.com/hardikbansal/gpt-enterprise-ui/logger"

func (adapter *DbAdapter) RunMigrations() {
	err := adapter.db.AutoMigrate(&User{})
	if err != nil {
		logger.LogProblem("Error auto migrating user table")
		return
	}
}
