package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/user"
	"github.com/muhammadaskar/casheer-be/app/user/usecase"
	"github.com/muhammadaskar/casheer-be/infrastructures/auth"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	authJwt     auth.JWTAuthentication
}

func NewUserHandler(userUseCase usecase.UserUseCase, auth auth.JWTAuthentication) *UserHandler {
	return &UserHandler{userUseCase, auth}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input user.RegisterInput
	message := "Register account failed"

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userUseCase.Register(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse(message, http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := customresponse.APIResponse(message, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "")
	response := customresponse.APIResponse("Account has been registered, admin will active your account", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}
		response := customresponse.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userUseCase.Login(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Login Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authJwt.GenerateToken(loggedInUser.ID, loggedInUser.Email, loggedInUser.Role)
	if err != nil {
		reponse := customresponse.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)
	response := customresponse.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}