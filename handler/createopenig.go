package handler

import (
	"net/http"

	"github.com/PedroPBSilva/gopportunities.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {

	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(context, "create-opening", opening)
}