package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/app/product/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(productUseCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase}
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	products, err := h.productUseCase.FindAll()
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := customresponse.APIResponse("List of products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input product.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	newProduct, err := h.productUseCase.Create(input)
	if err != nil {
		response := customresponse.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to create product", http.StatusCreated, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusCreated, response)

}
