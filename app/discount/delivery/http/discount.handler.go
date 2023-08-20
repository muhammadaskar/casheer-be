package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/discount"
	"github.com/muhammadaskar/casheer-be/app/discount/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type DiscountHandler struct {
	discountUseCase usecase.DiscountUseCase
}

func NewDiscountHandler(discountUseCase usecase.DiscountUseCase) *DiscountHandler {
	return &DiscountHandler{discountUseCase}
}

func (h *DiscountHandler) FindByID(c *gin.Context) {
	data, err := h.discountUseCase.FindByID()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := customresponse.APIResponse("Error to get discount", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Get discount", http.StatusOK, "success", discount.FormatDiscount(data))
	c.JSON(http.StatusOK, response)
}

func (h *DiscountHandler) Create(c *gin.Context) {
	var input discount.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create discount", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	newDiscount, err := h.discountUseCase.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to create discount", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to create new discount", http.StatusCreated, "success", discount.FormatDiscount(newDiscount))
	c.JSON(http.StatusCreated, response)
}

func (h *DiscountHandler) Update(c *gin.Context) {
	var input discount.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update discount", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	updateDiscount, err := h.discountUseCase.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to update discount", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to update new discount", http.StatusOK, "success", discount.FormatDiscount(updateDiscount))
	c.JSON(http.StatusOK, response)
}
