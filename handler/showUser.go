package handler

import (
	"github.com/Gabriel-Bitencort/AutoMailApi.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Show user
// @Description Show a user register
// @Tags Register
// @Accept json
// @Produce json
// @Param id query string true "User identification"
// @Success 200 {object} ShowUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /register [get]
func ShowUserHandler(ctx *gin.Context) {
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
	sendSuccess(ctx, "show-user", register)
}
