package handler

import (
	"github.com/Gabriel-Bitencort/AutoMailApi.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary List users
// @Description List all users
// @Tags Register
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [get]
func ListUsersHandler(ctx *gin.Context) {
	register := []schemas.UserRegister{}

	if err := db.Find(&register).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	sendSuccess(ctx, "list-users", register)
}
