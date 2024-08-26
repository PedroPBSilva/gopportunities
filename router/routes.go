package router

import (
	docs "github.com/PedroPBSilva/gopportunities.git/docs"
	"github.com/PedroPBSilva/gopportunities.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InicializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InicializeHandler()

	docs.SwaggerInfo.BasePath = "/api/v1"
	
	v1 := router.Group("/api/v1") 
	{
		v1.GET("/opening", handler.ShowOpeningHandler )

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	
	//Inicialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}