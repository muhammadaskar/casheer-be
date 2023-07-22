package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/auth"
	"github.com/muhammadaskar/casheer-be/app/helper"
	"github.com/muhammadaskar/casheer-be/app/user"
)

type UserHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *UserHandler {
	return &UserHandler{userService, authService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input user.RegisterInput
	message := "Register account failed"

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := helper.APIResponse(message, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "")
	response := helper.APIResponse("Account has been registered, admin will active your account", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID, loggedInUser.Email, loggedInUser.Role)
	if err != nil {
		reponse := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
