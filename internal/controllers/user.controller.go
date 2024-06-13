package controller

import (
	"GO-API-SERVER/internal/service"
	"GO-API-SERVER/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetByUserID(c *gin.Context) {
	response.ErrorResponse(c, 20003, "No need!")
}