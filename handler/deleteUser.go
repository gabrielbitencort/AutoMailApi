package handler

import (
	"fmt"
	"github.com/Gabriel-Bitencort/AutoMailApi.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Delete user
// @Description Delete a user register
// @Tags Register
// @Accept json
// @Produce json
// @Param id query string true "User identification"
// @Success 200 {object} DeleteUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /register [delete]
func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	register := schemas.UserRegister{}

	// Find User
	if err := db.First(&register, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("user with id: %s not found", id))
		return
	}

	// Delete User
	if err := db.Delete(&register).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-user", register)
}
