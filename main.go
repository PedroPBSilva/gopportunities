package main

import (
	"github.com/PedroPBSilva/gopportunities.git/config"
	"github.com/PedroPBSilva/gopportunities.git/router"
)

var (
	logger *config.Logger
)
func main() {
	//Instance Logger
	logger = config.GetLogger("main")
	
	//Inicialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config inicialization DB error: %v", err)
		return
	}

	//Inicialize router
	router.Inicialize()
}
