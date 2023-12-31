package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/user"
	"github.com/muhammadaskar/casheer-be/app/user/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
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
		errMessage := gin.H{"errors": "username or password doesnt match"}
		response := customresponse.APIResponse("Login Failed", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authJwt.GenerateToken(loggedInUser.ID, loggedInUser.Email, loggedInUser.Role)
	if err != nil {
		fmt.Println("masuk sini error generate token")
		reponse := customresponse.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, reponse)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)
	response := customresponse.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *UserHandler) UpdateNameOrEmail(c *gin.Context) {
	var inputData user.NameAndEmailInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Update user failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)

	updateUser, err := h.userUseCase.UpdateNameOrEmail(currentUser.ID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Update user failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataUser := gin.H{"name": updateUser.Name, "email": updateUser.Email}

	response := customresponse.APIResponse("Success to update user", http.StatusOK, "success", dataUser)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var inputData user.PasswordInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Update user failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)

	updateUser, err := h.userUseCase.UpdatePassword(currentUser.ID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Update user failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataUser := gin.H{"name": updateUser.Name, "email": updateUser.Email}

	response := customresponse.APIResponse("Success to update password", http.StatusOK, "success", dataUser)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserAdmin(c *gin.Context) {
	casheers, err := h.userUseCase.GetUserAdmin()
	if err != nil {
		response := customresponse.APIResponse("Failed to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get users", http.StatusOK, "success", casheers)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserCasheers(c *gin.Context) {
	casheers, err := h.userUseCase.GetUserCasheers()
	if err != nil {
		response := customresponse.APIResponse("Failed to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get users", http.StatusOK, "success", casheers)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUsersUnprocess(c *gin.Context) {
	casheers, err := h.userUseCase.GetUsersUnprocess()
	if err != nil {
		response := customresponse.APIResponse("Failed to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get users", http.StatusOK, "success", casheers)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUsersRejected(c *gin.Context) {
	casheers, err := h.userUseCase.GetUsersRejected()
	if err != nil {
		response := customresponse.APIResponse("Failed to get users", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get users", http.StatusOK, "success", casheers)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetTotalCasheer(c *gin.Context) {
	casheer, err := h.userUseCase.GetTotalCasheer()
	if err != nil {
		response := customresponse.APIResponse("Failed to get total casheer", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get total casheer", http.StatusOK, "success", casheer)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Activate(c *gin.Context) {
	var inputID user.GetUserIDInput
	err := c.BindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to accept user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userUseCase.Accept(inputID)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to accept user", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("User success to activated", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Reject(c *gin.Context) {
	var inputID user.GetUserIDInput
	err := c.BindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to reject user", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userUseCase.Reject(inputID)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to reject user", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("User success to rejected", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) ChangeToAdmin(c *gin.Context) {
	var inputID user.GetUserIDInput
	err := c.BindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to change role", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userUseCase.ChangeToAdmin(inputID)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to change role", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("User success to change role", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
