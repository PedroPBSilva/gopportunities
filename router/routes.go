package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InicializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1") 
	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET opening",
			})
		})

		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST opening",
			})
		})

		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "PUT opening",
			})
		})

		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "DELETE opening",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET openings",
			})
		})
	}
	
}