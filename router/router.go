package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//OBS: funções nos packages tem que iniciar com letra maiuscula senão nao consegue ser referenciada em outro local
func Inicialize() {
	fmt.Println("Starting API Rest!")
	//Inicialize Router
	router := gin.Default()
	
	//Inicialize Routes
	InicializeRoutes(router)

	//Run the server
	router.Run(":8080")
}