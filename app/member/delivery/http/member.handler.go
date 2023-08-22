package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/member"
	"github.com/muhammadaskar/casheer-be/app/member/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type MemberHandler struct {
	memberUseCase usecase.MemberUseCase
}

func NewMemberHandler(memberUseCase usecase.MemberUseCase) *MemberHandler {
	return &MemberHandler{memberUseCase}
}

func (h *MemberHandler) FindAll(c *gin.Context) {
	members, err := h.memberUseCase.FindAll()
	if err != nil {
		response := customresponse.APIResponse("Failed to get members", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get members", http.StatusOK, "success", members)
	c.JSON(http.StatusOK, response)
}

func (h *MemberHandler) Create(c *gin.Context) {
	var input member.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMesssage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create member", http.StatusUnprocessableEntity, "error", errorMesssage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	newMember, err := h.memberUseCase.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to create member", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to create member", http.StatusCreated, "success", newMember)
	c.JSON(http.StatusCreated, response)
}
