package handler

import (
	"fmt"
	"net/http"

	"github.com/PedroPBSilva/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		logger.Errorf("ID request is empty")
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameters").Error())
		return
	}

	opening := schemas.Opening{}

	// Find record from ID and populate variable opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("Opening with id:  %v not found", id)
		sendError(context, http.StatusNotFound, fmt.Sprintf("Opening with id:  %v not found", id))
		return
	}

	sendSucess(context, "show-opening", opening)
}