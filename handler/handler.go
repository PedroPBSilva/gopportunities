package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "Handler POST opening",
	})
}

func UpdateOpeningHandler(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "Handler PUT opening",
	})
}

func DeleteOpeningHandler(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "Handler DELETE opening",
	})
}

func ShowOpeningHandler(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "Handler SHOW opening",
	})
}

func ListOpeningsHandler(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "Handler LIST openings",
	})
}