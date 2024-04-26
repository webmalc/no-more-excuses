//go:build !test
// +build !test

package main

import (
	"webmalc/no-more-excuses/cmd"
	"webmalc/no-more-excuses/common/config"
	"webmalc/no-more-excuses/common/logger"
	"webmalc/no-more-excuses/internal/repositories"
	"webmalc/no-more-excuses/internal/server"
	"webmalc/no-more-excuses/internal/ui"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	appRepo := repositories.NewAppRepository(log)
	uiObj := ui.NewUI(appRepo)
	serverObj := server.NewServer(log)
	router := cmd.NewCommandRouter(
		log, serverObj, uiObj,
	)
	router.Run()
}
