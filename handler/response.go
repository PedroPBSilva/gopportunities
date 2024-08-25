package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, code int, message string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"status": "error",
		"message": message,
		"errorCode": code,
	})
}

func sendSucess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"status": "sucess",
		"message": fmt.Sprintf("Operation from handler: %s successfull", op),
		"data": data,
	})
}
