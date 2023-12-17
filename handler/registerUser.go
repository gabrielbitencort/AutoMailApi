package handler

import (
	"github.com/Gabriel-Bitencort/AutoMailApi.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Register user
// @Description Register a new user
// @Tags Register
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "Request body"
// @Success 200 {object} RegisterUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func RegisterUserHandler(ctx *gin.Context) {
	request := RegisterUserRequest{}

	_ = ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	register := schemas.UserRegister{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: request.PasswordHash,
	}

	if err := db.Create(&register).Error; err != nil {
		logger.Errorf("error registering user: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error registering user on database")
		return
	}

	sendSuccess(ctx, "register-user", register)
}
