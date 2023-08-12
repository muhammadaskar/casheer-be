package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/helper"
	"github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/domains"
)

type ProductHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *ProductHandler {
	return &ProductHandler{productService}
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input product.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	newProduct, err := h.productService.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create product", http.StatusCreated, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusCreated, response)

}
