package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/user_image/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type UserImageHandler struct {
	userImageUseCase usecase.UserImageUseCase
}

func NewUserImageHandler(userImageUseCase usecase.UserImageUseCase) *UserImageHandler {
	return &UserImageHandler{userImageUseCase}
}

func (h *UserImageHandler) FindByUserId(c *gin.Context) {
	currentUserID := c.MustGet("currentUser").(domains.User).ID

	userImage, err := h.userImageUseCase.FindByUserId(currentUserID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to get image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get image", http.StatusOK, "success", userImage)
	c.JSON(http.StatusOK, response)
}

func (h *UserImageHandler) CreateOrUpdate(c *gin.Context) {
	var input domains.UserImageInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to post image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.UserID = currentUser.ID

	newUserImage, err := h.userImageUseCase.CreateOrUpdate(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to post image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := customresponse.APIResponse("Success to post image", http.StatusOK, "success", newUserImage)
	c.JSON(http.StatusOK, response)
}
