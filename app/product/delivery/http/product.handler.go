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
	var query product.GetProductsQueryInput
	// page := c.DefaultQuery("page", "1") // Mengambil query parameter "page" dari URL, default: "1"

	err := c.ShouldBindQuery(&query)
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.productUseCase.FindAll(query)
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := customresponse.APIResponse("List of products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) FindById(c *gin.Context) {
	var input product.GetProductDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := customresponse.APIResponse("Error to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.productUseCase.FindById(input)
	if err != nil {
		response := customresponse.APIResponse("Error to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Product detail", http.StatusOK, "success", productDetail)
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
