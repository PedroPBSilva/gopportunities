package handler

import (
	"fmt"
	"net/http"

	"github.com/PedroPBSilva/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening key"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
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