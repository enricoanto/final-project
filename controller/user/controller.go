package user

import (
	"fmt"
	"net/http"

	"github.com/enricoanto/final-project/helper"
	model "github.com/enricoanto/final-project/repository"

	userService "github.com/enricoanto/final-project/service/user"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *userService.Service
}

func NewController(userService *userService.Service) *Controller {
	return &Controller{
		userService: userService,
	}
}

func (controller *Controller) Register(c *gin.Context) {
	var register RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(register); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	userRegister := model.User{
		FullName: register.FullName,
		Email:    register.Email,
		Password: register.Password,
		Balance:  register.Balance,
	}
	userRegister.Role = "customer"

	user, err := controller.userService.Register(userRegister)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := transformToRegisterResponse(user)

	helper.Success(c, http.StatusCreated, response)
}

func (controller *Controller) Login(c *gin.Context) {
	var login LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(login); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	userLogin := model.User{
		Email:    login.Email,
		Password: login.Password,
	}

	token, err := controller.userService.Login(userLogin)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	response := TokenResponse{
		Token: token,
	}

	helper.Success(c, http.StatusOK, response)
}

func (controller *Controller) UpdateBalance(c *gin.Context) {
	var updateBalance RegisterRequest

	if err := c.ShouldBindJSON(&updateBalance); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := helper.GetValidator().Struct(updateBalance); err != nil {
		helper.Error(c, http.StatusBadRequest, err)
		return
	}

	balance, err := controller.userService.UpdateBalance(2, updateBalance.Balance)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err)
		return
	}

	message := fmt.Sprintf("Your balance has been successfullly update to Rp %v", balance)

	response := MessageResponse{
		Message: message,
	}

	helper.Success(c, http.StatusOK, response)
}
