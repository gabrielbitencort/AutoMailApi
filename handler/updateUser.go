package handler

import (
	"github.com/Gabriel-Bitencort/AutoMailApi.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	register := schemas.UserRegister{}
	if err := db.First(&register, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	// Update user
	if request.Name != "" {
		register.Name = request.Name
	}
	if request.Email != "" {
		register.Email = request.Email
	}
	if request.PasswordHash != "" {
		register.PasswordHash = request.PasswordHash
	}

	// Save user
	if err := db.Save(&register).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, "update-user", register)
}
