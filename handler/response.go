package handler

import (
	"fmt"
	"github.com/Gabriel-Bitencort/AutoMailApi.git/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type RegisterUserResponse struct {
	Message string                       `json:"message"`
	Data    schemas.UserRegisterResponse `json:"data"`
}
type DeleteUserResponse struct {
	Message string                       `json:"message"`
	Data    schemas.UserRegisterResponse `json:"data"`
}
type ShowUserResponse struct {
	Message string                       `json:"message"`
	Data    schemas.UserRegisterResponse `json:"data"`
}
type ListUsersResponse struct {
	Message string                         `json:"message"`
	Data    []schemas.UserRegisterResponse `json:"data"`
}
type UpdateUserResponse struct {
	Message string                       `json:"message"`
	Data    schemas.UserRegisterResponse `json:"data"`
}
