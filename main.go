package main

import (
	"fmt"
	"github.com/hardikbansal/gpt-enterprise-ui/auth"
	"github.com/hardikbansal/gpt-enterprise-ui/config"
	"github.com/hardikbansal/gpt-enterprise-ui/core"
	"github.com/hardikbansal/gpt-enterprise-ui/db"
	"github.com/hardikbansal/gpt-enterprise-ui/handlers"
	"github.com/hardikbansal/gpt-enterprise-ui/llm"
	logger "github.com/hardikbansal/gpt-enterprise-ui/logger"
	"github.com/hardikbansal/gpt-enterprise-ui/server"
)

// main function
func main() {

	appConfig := config.GetInstance()
	if appConfig == nil {
		logger.LogPanic("config not found")
	}
	logger.InitiateLogger()

	dbService, err := db.GetDbAdapter("postgres://chatgpt:chatgpt@localhost:5432/chatgpt")
	authService := &auth.Service{}
	llmService := &llm.Service{}
	dbService.RunMigrations()
	if err != nil {
		panic("Db connection not working")
	}
	fmt.Println("Db connected")
	service := core.GetNewService(dbService, authService, llmService)
	ginHandler := handlers.NewApiHandler(service)
	server.StartServer(ginHandler, appConfig.Port)
}
