package main

import (
	"github.com/Gabriel-Bitencort/AutoMailApi.git/config"
	"github.com/Gabriel-Bitencort/AutoMailApi.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
