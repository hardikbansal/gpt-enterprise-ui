package main

import (
	"github.com/hardikbansal/gpt-enterprise-ui/config"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	"github.com/hardikbansal/gpt-enterprise-ui/handlers"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"github.com/hardikbansal/gpt-enterprise-ui/server"
)

// main function
func main() {

	appConfig := config.GetInstance()
	if appConfig == nil {
		logger.LogPanic("confif not found")
	}
	logger.InitiateLogger()

	service := core.GetNewService()
	ginHandler := handlers.NewApiHandler(service)
	server.StartServer(ginHandler, appConfig.Port)
}
